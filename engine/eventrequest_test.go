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

package engine

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cgrates/cgrates/config"
	"github.com/cgrates/cgrates/utils"
)

func TestEventRequestParseFieldDateTimeDaily(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*daily", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*daily", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "",
	}

	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*daily", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", expected, result)
	}
}

func TestEventReqParseFieldDateTimeTimeZone(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*daily", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*daily", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}

	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*daily", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestEventReqParseFieldDateTimeMonthly(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*monthly", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*monthly", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*monthly", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Month(), expected.Month()) {
		t.Errorf("Expected %v but received %v", finRes.Month(), expected.Month())
	}
}

func TestEventReqParseFieldDateTimeMonthlyEstimated(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*monthly_estimated", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*monthly_estimated", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*monthly_estimated", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Month(), expected.Month()) {
		t.Errorf("Expected %v but received %v", finRes.Month(), expected.Month())
	}
}

func TestEventReqParseFieldDateTimeYearly(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*yearly", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*yearly", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*yearly", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Year(), expected.Year()) {
		t.Errorf("Expected %v but received %v", finRes.Year(), expected.Year())
	}
}

func TestEventReqParseFieldDateTimeMetaUnlimited(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile(utils.MetaUnlimited, utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile(utils.MetaUnlimited, utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout(utils.MetaUnlimited, utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestEventReqParseFieldDateTimeEmpty(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestEventReqParseFieldDateTimeMonthEnd(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "Local",
	}
	result, err := EventReq.ParseField(fctTemp)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}

	expected, err := utils.ParseTimeDetectLayout("*month_endTest", utils.FirstNonEmpty(fctTemp.Timezone, config.CgrConfig().GeneralCfg().DefaultTimezone))
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	strRes := fmt.Sprintf("%v", result)
	finRes, err := time.Parse("“Mon Jan _2 15:04:05 2006”", strRes)
	if err != nil {
		t.Errorf("Expected %v but received %v", nil, err)
	}
	if !reflect.DeepEqual(finRes.Day(), expected.Day()) {
		t.Errorf("Expected %v but received %v", finRes.Day(), expected.Day())
	}
}

func TestAgentRequestParseFieldDateTimeError(t *testing.T) {
	tntTpl := config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep)
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, tntTpl, "", "", nil, nil)
	fctTemp := &config.FCTemplate{
		Type:     utils.MetaDateTime,
		Value:    config.NewRSRParsersMustCompile("*month_endTest", utils.InfieldSep),
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "/",
	}
	_, err := EventReq.ParseField(fctTemp)
	expected := "time: invalid location name"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected <%+v> but received <%+v>", expected, err)
	}
}

func TestEventReqParseFieldDateTimeError2(t *testing.T) {
	prsr, err := config.NewRSRParsersFromSlice([]string{"2.", "~*req.CGRID<~*opts.Converter>"})
	if err != nil {
		t.Fatal(err)
	}
	EventReq := NewEventRequest(utils.MapStorage{}, nil, nil, prsr, "", "", nil, nil)
	fctTemp := &config.FCTemplate{Type: utils.MetaDateTime,
		Value:    prsr,
		Layout:   "“Mon Jan _2 15:04:05 2006”",
		Timezone: "/",
	}

	_, err = EventReq.ParseField(fctTemp)
	expected := utils.ErrNotFound
	if err == nil || err != expected {
		t.Errorf("Expected <%+v> but received <%+v>", expected, err)
	}
}