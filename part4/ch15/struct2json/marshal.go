package main

import (
	"bytes"
	"reflect"
)

// MarshalJSON ...
func MarshalJSON(v interface{}) (string, error) {
	// 准备一个缓冲
	var b bytes.Buffer
	var err error

	// 将任意值转换为json并输出到缓冲
	if err = writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	}

	return "", err
}
