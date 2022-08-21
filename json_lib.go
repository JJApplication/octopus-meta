/*
Create: 2022/7/22
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

// json
import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

// 使用autoEnv时 默认使用jsoniter
// 否则使用encoding/json

var jsonAPI jsoniter.API
var lock sync.Once

type OctopusJSON interface {
	Marshal(v interface{}) ([]byte, error)
	MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	UnmarshalFromString(str string, v interface{}) error
}

// 一个单例的被任意调用的API
// 返回一个可用的OctopusJSON 接口
func json(autoEnv bool) OctopusJSON {
	if autoEnv {
		jsonAPI = jsoniter.Config{}.Froze()
		lock.Do(func() {
			register(autoEnv)
		})
		return jsonAPI
	}
	return OriginJSON
}

func register(autoEnv bool) {
	jsoniter.RegisterTypeDecoderFunc("string", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		// 进行值拷贝
		cp := iter.Read()
		if !autoEnv {
			*((*string)(ptr)) = cp.(string)
			return
		}
		if checkValFromEnv(cp) {
			enVal := getValFromEnv(cp.(string))
			if enVal == "" {
				*((*string)(ptr)) = ""
			} else {
				*((*string)(ptr)) = enVal
			}
		} else {
			*((*string)(ptr)) = cp.(string)
		}
	})

	jsoniter.RegisterTypeDecoderFunc("int64", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		cp := iter.Read()
		if !autoEnv {
			cpo, ok := cp.(float64)
			if !ok {
				// string无法解析
				*((*int64)(ptr)) = 0
			} else {
				*((*int64)(ptr)) = int64(math.Ceil(cpo))
			}

			return
		}
		if checkValFromEnv(cp) {
			enVal := getValFromEnv(cp.(string))
			if enVal == "" {
				*((*int64)(ptr)) = 0
			} else {
				b, _ := strconv.Atoi(enVal)
				*((*int64)(ptr)) = int64(b)
			}
		} else {
			*((*int64)(ptr)) = cp.(int64)
		}
	})

	jsoniter.RegisterTypeDecoderFunc("int", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		cp := iter.Read()
		if !autoEnv {
			cpo, ok := cp.(float64)
			if !ok {
				// string无法解析
				*((*int)(ptr)) = 0
			} else {
				*((*int)(ptr)) = int(math.Ceil(cpo))
			}
			return
		}
		if checkValFromEnv(cp) {
			enVal := getValFromEnv(cp.(string))
			if enVal == "" {
				*((*int)(ptr)) = 0
			} else {
				b, _ := strconv.Atoi(enVal)
				*((*int)(ptr)) = b
			}
		} else {
			*((*int)(ptr)) = int(math.Ceil(cp.(float64)))
		}
	})

	jsoniter.RegisterTypeDecoderFunc("bool", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		cp := iter.Read()
		if checkValFromEnv(cp) {
			enVal := getValFromEnv(cp.(string))
			if enVal == "" || enVal == "false" || enVal == "False" {
				*((*bool)(ptr)) = false
			} else if enVal == "yes" || enVal == "Yes" {
				*((*bool)(ptr)) = true
			} else {
				*((*bool)(ptr)) = false
			}
		} else {
			*((*bool)(ptr)) = cp.(bool)
		}
	})
}

func checkValFromEnv(v interface{}) bool {
	val, ok := v.(string)
	if ok {
		return strings.HasPrefix(val, "$")
	}
	return false
}

func getValFromEnv(v string) string {
	realVal := strings.TrimPrefix(v, "$")
	return os.Getenv(realVal)
}

type ex struct {
	jsoniter.DummyExtension
}

// UpdateStructDescriptor 更新字段
func (e ex) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, f := range structDescriptor.Fields {
		f.Decoder = &decodeFunc{func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.WhatIsNext() != jsoniter.NilValue {
				if iter.Error == nil {
					iter.Error = errors.New("")
				}
			} else {
				iter.Skip()
			}
		}}

		f.Encoder = &encodeFunc{func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			if ptr == nil {
				stream.WriteNil()
			} else if stream.Error == nil {
				stream.Error = errors.New("ss")
			}
		}, nil}
	}
}

// decoder
type decodeFunc struct {
	fun jsoniter.DecoderFunc
}

func (d *decodeFunc) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	d.fun(ptr, iter)
}

type encodeFunc struct {
	fun         jsoniter.EncoderFunc
	isEmptyFunc func(ptr unsafe.Pointer) bool
}

func (e *encodeFunc) IsEmpty(ptr unsafe.Pointer) bool {
	if e.isEmptyFunc == nil {
		return false
	}
	return e.isEmptyFunc(ptr)
}

func (e *encodeFunc) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e.fun(ptr, stream)
}
