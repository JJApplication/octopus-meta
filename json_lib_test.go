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

type T struct {
	Name string
	Age  int
	Meta TT
}

type TT struct {
	Sp string
}

func TestJson(t *testing.T) {
	var res T
	err := json(false).UnmarshalFromString(`{"name": "test", "age": 10}`, &res)
	if assert.NoError(t, err) {
		t.Log("test json success")
	} else {
		t.Error("test json failed")
	}
	t.Logf("%+v", res)
}

func TestJsonWithEnv(t *testing.T) {
	var res T
	os.Setenv("test", "jsonTest")
	os.Setenv("age", "10")
	err := json(true).UnmarshalFromString(`{"name": "$test", "age": "$age"}`, &res)
	if assert.NoError(t, err) && assert.Equal(t, T{
		Name: "jsonTest",
		Age:  10,
		Meta: TT{},
	}, res) {
		t.Log("test envFlag true success")
	} else {
		t.Error("test envFlag true failed")
	}
	t.Logf("%+v", res)
}

func TestJsonWithNoEnv(t *testing.T) {
	var res T
	os.Setenv("test", "jsonTest")
	os.Setenv("age", "10")
	err := json(false).UnmarshalFromString(`{"name": "$test", "age": "$age"}`, &res)
	if assert.Error(t, err) {
		t.Log("test envFlag false success")
	} else {
		t.Error("test envFlag false failed")
	}
	t.Logf("%+v", res)
}
