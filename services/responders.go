/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package services

import (
	"fmt"
	"sync"

	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/servmanager"
	"github.com/cgrates/cgrates/utils"
	"github.com/cgrates/rpcclient"
)

// NewResponderService returns the Resonder Service
func NewResponderService() *ResponderService {
	return &ResponderService{
		connChan: make(chan rpcclient.RpcClientConnection, 1),
	}
}

// ResponderService implements Service interface
type ResponderService struct {
	sync.RWMutex
	resp     *engine.Responder
	connChan chan rpcclient.RpcClientConnection
}

// Start should handle the sercive start
// For this service the start should be called from RAL Service
func (resp *ResponderService) Start(sp servmanager.ServiceProvider, waitCache bool) (err error) {
	if resp.IsRunning() {
		return fmt.Errorf("service aleady running")
	}

	var thdS, stats rpcclient.RpcClientConnection
	if thdS, err = sp.NewConnection(utils.ThresholdS, sp.GetConfig().RalsCfg().ThresholdSConns); err != nil {
		utils.Logger.Crit(fmt.Sprintf("<%s> Could not connect to %s, error: %s",
			utils.RALService, utils.ThresholdS, err.Error()))
		return
	}
	if stats, err = sp.NewConnection(utils.StatS, sp.GetConfig().RalsCfg().StatSConns); err != nil {
		utils.Logger.Crit(fmt.Sprintf("<%s> Could not connect to %s, error: %s",
			utils.RALService, utils.StatS, err.Error()))
		return
	}
	if thdS != nil {
		engine.SetThresholdS(thdS) // temporary architectural fix until we will have separate AccountS
	}
	if stats != nil {
		engine.SetStatS(stats)
	}
	resp.Lock()
	defer resp.Unlock()
	resp.resp = &engine.Responder{
		ExitChan:         sp.GetExitChan(),
		MaxComputedUsage: sp.GetConfig().RalsCfg().MaxComputedUsage,
	}

	if !sp.GetConfig().DispatcherSCfg().Enabled {
		sp.GetServer().RpcRegister(resp.resp)
	}

	utils.RegisterRpcParams("", resp.resp)

	resp.connChan <- resp.resp // Rater done
	return
}

// GetIntenternalChan returns the internal connection chanel
func (resp *ResponderService) GetIntenternalChan() (conn chan rpcclient.RpcClientConnection) {
	return resp.connChan
}

// Reload handles the change of config
func (resp *ResponderService) Reload(sp servmanager.ServiceProvider) (err error) {
	var thdS, stats rpcclient.RpcClientConnection
	if thdS, err = sp.NewConnection(utils.ThresholdS, sp.GetConfig().RalsCfg().ThresholdSConns); err != nil {
		utils.Logger.Crit(fmt.Sprintf("<%s> Could not connect to %s, error: %s",
			utils.RALService, utils.ThresholdS, err.Error()))
		return
	}
	if stats, err = sp.NewConnection(utils.StatS, sp.GetConfig().RalsCfg().StatSConns); err != nil {
		utils.Logger.Crit(fmt.Sprintf("<%s> Could not connect to %s, error: %s",
			utils.RALService, utils.StatS, err.Error()))
		return
	}
	resp.Lock()
	resp.resp.MaxComputedUsage = sp.GetConfig().RalsCfg().MaxComputedUsage // this may cause concurrency problems
	if thdS != nil {
		engine.SetThresholdS(thdS) // temporary architectural fix until we will have separate AccountS
	}
	if stats != nil {
		engine.SetStatS(stats)
	}
	resp.Unlock()
	return
}

// Shutdown stops the service
func (resp *ResponderService) Shutdown() (err error) {
	resp.Lock()
	resp.resp = nil
	<-resp.connChan
	resp.Unlock()
	return
}

// GetRPCInterface returns the interface to register for server
func (resp *ResponderService) GetRPCInterface() interface{} {
	return resp.resp
}

// IsRunning returns if the service is running
func (resp *ResponderService) IsRunning() bool {
	resp.RLock()
	defer resp.RUnlock()
	return resp != nil && resp.resp != nil
}

// ServiceName returns the service name
func (resp *ResponderService) ServiceName() string {
	return utils.ResponderS
}

// GetResponder returns the responder created
func (resp *ResponderService) GetResponder() *engine.Responder {
	resp.RLock()
	defer resp.RUnlock()
	return resp.resp
}