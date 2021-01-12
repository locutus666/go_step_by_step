package main

import (
	"bytes"
	"reflect"
)

// MarshalJSON 序列化主函数
// MarshalJson是序列化过程的主要函数入口，通过这个函数会调用不同类型的子序列化函数writeAny
func MarshalJSON(v interface{}) (string, error) { // 传入空接口interface{}数据，并将这个数据转换为JSON字符串返回，如果发生错误，则返回错误信息。
	// 使用bytes.Buffer构建一个缓冲，在大量字符串连接时，推荐使用这个结构
	var b bytes.Buffer
	var err error

	// 调用writeAny函数，将bytes.Buffer（变量b）以指针的方式传入，以方便将各种类型的数据都写入到这个bytes.Buffer中
	// 同时，把空接口v转换为反射值对象，并作为参数传入
	if err = writeAny(&b, reflect.ValueOf(v)); err == nil { // 如果err == nil，就是没有错误发生
		return b.String(), nil // 将 bytes.Buffer 的内容转换为字符串，井返回
	}

	return "", err // 如果err != nil，就是发生了错误，返回空字符串结果和错误
}

/*
MarshalJSON这个函数，其实是对writeAny函数的一个封装，把外部的空接口类型interface{}转换为内部的reflect.Value类型，同时构建输出缓冲，将一些复杂的操作简化，方便外部使用。
*/
