/*
Create: 2022/8/25
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

// Alarm 告警信息模型
// 在使用时 存储数据库时增加mongo model
type Alarm struct {
	Title        string `json:"title" bson:"title"`
	Level        string `json:"level" bson:"level"`
	Message      string `json:"message" bson:"message"`
	AlarmExtends `json:",omitempty" bson:",inline"`
}

// AlarmExtends 告警扩展信息
type AlarmExtends struct {
	Source string `json:"source" bson:"source"` // 告警来源
	User   string `json:"user" bson:"user"`     // 操作用户
}

const (
	SourceOther     = "other"
	SourceApp       = "app"
	SourceSys       = "system"
	SourceContainer = "container"
)
