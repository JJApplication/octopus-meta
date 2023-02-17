/*
Package octopus_meta
用于解析octopus结构的库
数据结构转化 json->json go->json
用于从原生go转为meta数据
*/

package octopus_meta

import (
	"os"
	"path"
	"strings"
)

const (
	_version   = "1.0.0"
	_copyright = "projectJ"
	_uri       = "renj.io"
)

// Info 返回版本信息
func Info() (string, string, string) {
	return _version, _copyright, _uri
}

// Load 从指定路径加载模型
// 调用SetOctopusMetaDir来设置全局使用的路径
func Load(p string) (map[string]App, error) {
	if OctopusMetaDir != "" {
		p = OctopusMetaDir
	}
	if p == "" {
		return nil, ErrMetaDir
	}
	if isDirExist(p) {
		return load(p)
	}
	return nil, ErrMetaDir
}

// AutoLoad 从环境变量APP_ROOT下规定的路径自动加载
func AutoLoad() (map[string]App, error) {
	autoLoadDir := path.Join(os.Getenv("APP_ROOT"), MetaAutoLoadDir)
	autoLoadDirNoHide := path.Join(os.Getenv("APP_ROOT"), MetaAutoLoadDirNoHide)
	if isDirExist(autoLoadDir) {
		return load(autoLoadDir)
	}
	if isDirExist(autoLoadDirNoHide) {
		return load(autoLoadDirNoHide)
	}

	return nil, ErrMetaDir
}

// LoadApp 加载指定的app
// 仅在配置全局路径时生效
func LoadApp(appName string) (App, error) {
	return loadAPP(appName)
}

// NewMetaDir 创建新的meta目录
// 如果创建成功 默认会使用此目录作为全局目录
func NewMetaDir(p string) error {
	e := newDir(p)
	if e != nil {
		return e
	}
	SetOctopusMetaDir(p)
	return nil
}

// NewAppMeta 创建新的空app meta
// 自动加载目录和全局目录都不存在时 调用SetOctopusMetaDir来设置全局使用的路径
func NewAppMeta(appName string) error {
	return newApp(appName)
}

// NewApp 返回一个默认的新APP
func NewApp(appName string) App {
	app := DefaultApp
	app.Name = appName
	app.ID = "app_" + strings.ToLower(appName)
	return app
}

// DelAppMeta 删除app
func DelAppMeta(appName string) error {
	return delApp(appName)
}

// SaveAppMeta 保存app meta
func SaveAppMeta(app App, appName string) error {
	return saveApp(app, appName)
}
