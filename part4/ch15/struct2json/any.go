package main

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

// 下面例子只支持整型、字符串、切片、结构体类型序列化为JSON格式。如果需要扩充类型，可以在writeAny函数中添加对应类型
// writeAny函数传入一个字节缓冲和反射值对象，将反射值对象转换为JSON文本格式并写入字节缓冲中
func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	// 根据传入反射值对象的种类进行判断，如字符串、整型、切片及结构体
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String())) // reflect.Value类型调用String函数，将传入值转换为字符串；strconv.Quote函数提供了比较正规的封装，将字符串用双引号括起来
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10)) // 把整型数据以十进制格式，使用strconv.Formatlnt函数格式化为字符串；使用bytes.Buffer的WriteString函数把字符串写入字节缓冲中
	// 使用writeSlice函数，把切片序列化为JSON
	case reflect.Slice:
		return writeSlice(buff, value)
	// 使用writeStruct函数，把切片序列化为JSON
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupport kind :" + value.Kind().String()) // 遇到不能识别的种类，就返回错误
	}

	return nil
}

// writeAny函数是整个序列化中非常重要的环节，可以通过扩充switch中的种类，扩充序列化能识别的类型。
