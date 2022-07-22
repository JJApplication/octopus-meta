/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"os"
	"testing"
)

// 测试

func TestLoad(t *testing.T) {
	apps, err := Load("./test")
	if err != nil {
		t.Errorf("err load apps %v\n", err)
		t.SkipNow()
	}
	t.Log(apps)
}

func TestAutoLoad(t *testing.T) {
	os.Setenv("APP_ROOT", "./test")
	apps, err := AutoLoad()
	if err != nil {
		t.Errorf("err load apps %v\n", err)
		t.SkipNow()
	}
	t.Log(apps)
}

func TestLoadApp(t *testing.T) {
	SetOctopusMetaDir("./test")
	app, err := loadAPP("test")
	if err != nil {
		t.Errorf("err load apps %v\n", err)
		t.SkipNow()
	}
	t.Log(app)
}

func TestNewMetaDir(t *testing.T) {
	err := NewMetaDir("tmp")
	if err != nil {
		t.Errorf("err new meta dir %v\n", err)
		t.SkipNow()
	}
	if isDirExist("tmp") {
		t.Log("success")
		t.Log("clear tmp dir")
		os.RemoveAll("tmp")
	}
}

func TestNewAppMeta(t *testing.T) {
	SetOctopusMetaDir("./test")
	err := NewAppMeta("test1")
	if err != nil {
		t.Errorf("err new meta %v\n", err)
		t.SkipNow()
	}
}

func TestDelAppMeta(t *testing.T) {
	SetOctopusMetaDir("./test")
	err := DelAppMeta("test1")
	if err != nil {
		t.Errorf("err del meta %v\n", err)
		t.SkipNow()
	}
}

type MetaModel struct {
	ID string
}

type TestModel struct {
	Name string
	Bool bool
	Int  int
	M    MetaModel
}

func TestAutoEnv(t *testing.T) {
	AutoEnv()
	os.Setenv("T", "test")
	var app TestModel
	err := json(true).Unmarshal([]byte(`{"Name":"$T", "Bool": "$T"}`), &app)
	t.Log(err)
	t.Log(app)
	modifyAppEnv(&app)
	t.Log(app)
}
