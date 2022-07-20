/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

const (
	MetaAutoLoadDir       = ".octopus"
	MetaAutoLoadDirNoHide = "octopusMeta"
	MetaSuffix            = ".pig"
)

const (
	TypeService    = "Service"    // 服务
	TypeWebFront   = "FrontEnd"   // 前端
	TypeMiddleWare = "MiddleWare" // 中间件
	TypeDataStore  = "DataStore"  // 数据层
)

// app的发布状态
const (
	Published = "published" // 已发布
	Testing   = "testing"   // 测试中
	Pending   = "pending"   // 待发布
)

// 配置文件类型
const (
	ConfNginx    = "nginx"
	ConfGunicorn = "gunicorn"
)
