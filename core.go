/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/landers1037/configen"
)

var OctopusMetaDir = ""
var autoEnv bool

// SetOctopusMetaDir 设置配置需要读取的默认路径
// 后续所有的读取都会走这个路径 autoload时不生效
func SetOctopusMetaDir(p string) {
	OctopusMetaDir = p
}

// AutoEnv 自动替换值为$val的变量为环境变量val
func AutoEnv() {
	autoEnv = true
}

// 读取meta下的app目录 p目录经过计算一定存在
func getAPPCfs(p string) ([]string, error) {
	var cfs []string
	if !isDirExist(p) {
		return cfs, ErrMetaDir
	}

	err := filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			if !d.IsDir() {
				cfs = append(cfs, path)
			}
		}

		return err
	})

	if err != nil {
		return nil, ErrWalkMetaDir
	}

	return cfs, nil
}

// 加载全部app信息
func loadAllCfs(cfs []string) (map[string]App, bool) {
	var loadStatus = true // 记录是否有app配置加载失败
	tm := make(map[string]App, len(cfs))
	for _, c := range cfs {
		var appCfg App
		err := configen.ParseConfig(&appCfg, configen.Pig, c)
		if err != nil || reflect.DeepEqual(appCfg, App{}) {
			loadStatus = false
			continue
		}

		// get name
		name := strings.TrimSuffix(filepath.Base(c), filepath.Ext(c))
		// save to map
		tm[name] = appCfg
	}

	return tm, loadStatus
}

// 加载单独app 不区分大小写
// 只工作在OctopusMetaDir存在或者自动加载生效时
func loadAPP(name string) (App, error) {
	p := chooseDir()
	if p == "" {
		return App{}, ErrMetaDir
	}

	var cfs []string
	var cf string
	name = strings.ToLower(name)
	_ = filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			if !d.IsDir() {
				cfs = append(cfs, path)
			}
		}

		return err
	})
	for _, v := range cfs {
		if strings.ToLower(strings.TrimSuffix(filepath.Base(v), filepath.Ext(v))) == name {
			cf = v
		}
	}
	if cf == "" {
		return App{}, ErrAPP
	}

	// 加载
	var appCfg App
	err := configen.ParseConfig(&appCfg, configen.Pig, cf)
	return appCfg, err
}

// 自动提取环境变量的数据
func autoSetEnv(app *App) {

}

// 新建meta路径
func newDir(p string) error {
	return os.MkdirAll(p, 0644)
}

// 新建空app文件
func newApp(appName string) error {
	p := chooseDir()
	if p == "" {
		return ErrMetaDir
	}
	return configen.SaveConfig(
		App{
			Name:          appName,
			ID:            "app_" + strings.ToLower(appName),
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
				Domain:      "",
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
		},
		configen.Pig,
		path.Join(
			p, appName+MetaSuffix))
}

// 删除app
func delApp(appName string) error {
	p := chooseDir()
	if p == "" {
		return ErrMetaDir
	}
	f := path.Join(p, appName+MetaSuffix)
	if isDirExist(f) {
		l := sync.Mutex{}
		l.Lock()
		defer l.Unlock()
		err := os.RemoveAll(f)
		if err != nil {
			return err
		}
		return nil
	}
	return ErrMetaDir
}

func saveApp(app App, appName string) error {
	p := chooseDir()
	if p == "" {
		return ErrMetaDir
	}
	return configen.SaveConfig(
		app,
		configen.Pig,
		path.Join(
			p, appName+MetaSuffix))
}
