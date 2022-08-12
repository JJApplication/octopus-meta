/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"github.com/gookit/validate"
)

// App model for app
type App struct {
	Name          string `json:"name" validate:"required" bson:"name"` // 服务名称
	ID            string `json:"id" validate:"required" bson:"id"`     // 服务唯一ID
	Type          string `json:"type" bson:"type"`                     // service | middleware
	ReleaseStatus string `json:"release_status" bson:"release_status"` // published | pending | testing
	EngDes        string `json:"eng_des" bson:"eng_des"`               // 英文描述
	CHSDes        string `json:"chs_des" bson:"chs_des"`               // 中文描述
	Link          string `json:"link" bson:"link"`                     // 服务对外提供的URL链接
	// 管理项
	ManageCMD CMD `json:"manage_cmd" bson:"manage_cmd"` // 管理命令组
	// 元数据
	Meta Meta `json:"meta" bson:"meta"` // 服务元数据
	// 动态依赖配置
	RunData RunData `json:"run_data" bson:"run_data"` // 服务运行时数据

	ResourceLimit ResourceLimit `json:"resource_limit" bson:"resource_limit"` // 运行时资源限制
}

// CMD 服务的管理脚本
type CMD struct {
	Start     string `json:"start" bson:"start"`
	Stop      string `json:"stop" bson:"stop"`
	Restart   string `json:"restart" bson:"restart"`
	ForceKill string `json:"force_kill" bson:"force_kill"`
	Check     string `json:"check" bson:"check"`
}

// Meta 服务的元数据
type Meta struct {
	Author      string   `json:"author" bson:"author"`
	Domain      string   `json:"domain" bson:"domain"`
	Language    []string `json:"language" bson:"language"`
	CreateDate  string   `json:"create_date" bson:"create_date"`
	Version     string   `json:"version" bson:"version"`
	DynamicConf bool     `json:"dynamic_conf" bson:"dynamic_conf"` // 是否需要生成配置文件
	ConfType    string   `json:"conf_type" bson:"conf_type"`       // nginx | gunicorn
	ConfPath    string   `json:"conf_path" bson:"conf_path"`       // 支持绝对和相对路径
}

// RunData 运行时依赖
type RunData struct {
	Envs       []string `json:"envs" bson:"envs"` // just like `Name=Diri`
	Ports      []int    `json:"ports" bson:"ports"`
	RandomPort bool     `json:"random_port" bson:"random_port"` // if using random port
	Host       string   `json:"host" bson:"host"`               // always must be localhost
}

// ResourceLimit 运行时资源限制
type ResourceLimit struct {
	MinCpu     int `json:"min_cpu" bson:"min_cpu"`           // 最小cpu使用率 %
	MaxCpu     int `json:"max_cpu" bson:"max_cpu"`           // 最大cpu使用率 %
	MinMem     int `json:"min_mem" bson:"min_mem"`           // 最小内存使用率 bytes
	MaxMem     int `json:"max_mem" bson:"max_mem"`           // 最大内存使用率 bytes
	AveCpuPeak int `json:"ave_cpu_peak" bson:"ave_cpu_peak"` // 稳定负载下平均cpu峰值
	AveMemPeak int `json:"ave_mem_peak" bson:"ave_mem_peak"` // 稳定负载下平均mem峰值
	MaxRead    int `json:"max_read" bson:"max_read"`         // 最大读取速率
	MaxWrite   int `json:"max_write" bson:"max_write"`       // 最大写入速率
	MaxRequest int `json:"max_request" bson:"max_request"`   // 单位时间最大请求数
	MaxClient  int `json:"max_client" bson:"max_client"`     // 单位时间最大客户端连接数
}

// MetaData 导出App别名
type MetaData App

// Validate 适用于model的检查器
func (app *App) Validate() bool {
	if !validate.Struct(app).Validate() {
		return false
	}

	return true
}
