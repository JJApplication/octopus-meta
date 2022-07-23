/*
Create: 2022/7/22
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 类型测试

func TestOctopus_ParseString(t *testing.T) {
	var app App
	err := OctopusIterator.ParseString(&app, `{"name": "testApp"}`)
	if assert.NoError(t, err) && assert.Equal(t, App{Name: "testApp"}, app) {
		t.Log("test Octopus_ParseString success")
	} else {
		t.Log("test Octopus_ParseString failed")
	}
	t.Log(app)
}

func TestOctopus_ParseStringEnv(t *testing.T) {
	var app App
	os.Setenv("name", "envApp")
	err := OctopusIterator.ParseString(&app, `{"name": "$name"}`)
	if assert.NoError(t, err) && assert.Equal(t, App{Name: "envApp"}, app) {
		t.Log("test Octopus_ParseStringEnv success")
	} else {
		t.Log("test Octopus_ParseStringEnv failed")
	}
	t.Log(app)
}

func TestOctopus_ParseFile(t *testing.T) {
	var app App
	os.Setenv("name", "envApp")
	err := OctopusIterator.Parse(&app, "./test/test.pig")
	if assert.NoError(t, err) && assert.Equal(t,
		App{
			Name:          "envApp",
			ID:            "app_t",
			Type:          TypeService,
			ReleaseStatus: Published,
			EngDes:        "default english description",
			CHSDes:        "默认中文描述",
			ManageCMD: CMD{
				Start:     "start.sh",
				Stop:      "stop.sh",
				Restart:   "restart.sh",
				ForceKill: "kill.sh",
				Check:     "check.sh",
			},
			Meta: Meta{
				Author:      "",
				Domain:      "renj.io",
				Language:    []string{},
				CreateDate:  "",
				Version:     "1.0.0",
				DynamicConf: false,
				ConfType:    "",
				ConfPath:    "",
			},
			RunData: RunData{
				Envs:       []string{},
				Ports:      []int{},
				RandomPort: true,
				Host:       "localhost",
			},
		}, app) {
		t.Log("test Octopus_ParseFileEnv success")
	} else {
		t.Log("test Octopus_ParseFileEnv failed")
	}
	t.Log(app)
}

func TestOctopus_ParseFileWithNoEnv(t *testing.T) {
	var app App
	os.Setenv("name", "envApp")
	o := Octopus{AutoEnv: false}
	err := o.Parse(&app, "./test/test.pig")
	if assert.NoError(t, err) && assert.Equal(t,
		App{
			Name:          "$name",
			ID:            "app_t",
			Type:          TypeService,
			ReleaseStatus: Published,
			EngDes:        "default english description",
			CHSDes:        "默认中文描述",
			ManageCMD: CMD{
				Start:     "start.sh",
				Stop:      "stop.sh",
				Restart:   "restart.sh",
				ForceKill: "kill.sh",
				Check:     "check.sh",
			},
			Meta: Meta{
				Author:      "",
				Domain:      "renj.io",
				Language:    []string{},
				CreateDate:  "",
				Version:     "1.0.0",
				DynamicConf: false,
				ConfType:    "",
				ConfPath:    "",
			},
			RunData: RunData{
				Envs:       []string{},
				Ports:      []int{},
				RandomPort: true,
				Host:       "localhost",
			},
		}, app) {
		t.Log("test Octopus_ParseFileEnv success")
	} else {
		t.Log("test Octopus_ParseFileEnv failed")
	}
	t.Log(app)
}

func TestOctopus_ToString(t *testing.T) {
	var app App
	os.Setenv("name", "envApp")
	o := Octopus{AutoEnv: true}
	str, err := o.ToString(&app, "./test/test.pig")
	if assert.NoError(t, err) && assert.Equal(t, `{"name":"","id":"","type":"","release_status":"","eng_des":"","chs_des":"","manage_cmd":{"start":"","stop":"","restart":"","force_kill":"","check":""},"meta":{"author":"","domain":"","language":null,"create_date":"","version":"","dynamic_conf":false,"conf_type":"","conf_path":""},"run_data":{"envs":null,"ports":null,"random_port":false,"host":""}}`, str) {
		t.Log("test Octopus_ToJson success")
	} else {
		t.Log("test Octopus_ToJson failed")
	}
}

func TestOctopus_Save(t *testing.T) {
	var app App
	o := Octopus{AutoEnv: true}
	err := o.Save(&app, "./test/testSave.pig")
	var actualApp App
	err = o.Parse(&actualApp, "./test/testSave.pig")
	if assert.NoError(t, err) && assert.Equal(t, app, actualApp) {
		t.Log("test Octopus_Save success")
	} else {
		t.Log("test Octopus_Save failed")
	}
	defer os.RemoveAll("./test/testSave.pig")
}
