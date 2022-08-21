/*
Create: 2022/8/22
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	js "encoding/json"
)

type originJSON struct{}

// OriginJSON 使用go encoding/json
var OriginJSON = originJSON{}

func (o originJSON) MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return js.MarshalIndent(v, prefix, indent)
}

func (o originJSON) UnmarshalFromString(str string, v interface{}) error {
	return js.Unmarshal([]byte(str), v)
}

func (o originJSON) Marshal(v interface{}) ([]byte, error) {
	return js.Marshal(v)
}

func (o originJSON) Unmarshal(data []byte, v interface{}) error {
	return js.Unmarshal(data, v)
}
