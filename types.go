/*
Create: 2022/7/22
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"io/ioutil"
	"reflect"
	"sync"
)

// 自定义类型 用于反序列化和序列化

// Octopus 初始化或使用全局定义的octopus
type Octopus struct {
	Type    string
	AutoEnv bool
}

var OctopusIterator = Octopus{Type: "default", AutoEnv: true}

func (o *Octopus) j() OctopusJSON {
	return json(o.AutoEnv)
}

// Parse 从meta文件中解析
// v interface 必须是一个结构体指针
func (o *Octopus) Parse(v interface{}, f string) error {
	if !o.checkPtr(v) {
		return ErrNotPtr
	}
	data, err := o.ReadFile(f)
	if err != nil {
		return err
	}
	return o.j().Unmarshal(data, v)
}

// ParseString 从字符串中解析
func (o *Octopus) ParseString(v interface{}, s string) error {
	if !o.checkPtr(v) {
		return ErrNotPtr
	}
	return o.j().UnmarshalFromString(s, v)
}

// Save 保存数据到文件
func (o *Octopus) Save(v interface{}, f string) error {
	b, e := o.j().MarshalIndent(v, "", "  ")
	if e != nil {
		return e
	}
	lock := sync.Mutex{}
	lock.Lock()

	defer lock.Unlock()
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ToString 转为字符串
func (o *Octopus) ToString(v interface{}, f string) (string, error) {
	b, e := o.j().Marshal(v)
	return string(b), e
}

// ReadFile 读取文件内容
func (o *Octopus) ReadFile(f string) ([]byte, error) {
	lock := sync.Mutex{}
	lock.Lock()
	data, err := ioutil.ReadFile(f)
	defer lock.Unlock()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (o *Octopus) checkPtr(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}
