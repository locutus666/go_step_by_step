package main

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

// 下面例子只支持整型、字符串、切片、结构体类型序列化为JSON格式。如果需要扩充类型，可以在writeAny函数中添加对应类型
func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupport kind :" + value.Kind().String())
	}

	return nil
}

/*
以上的序列化方式，在程序功能和结构上还有一些不足：

没有处理各种异常情况，切片或结构体为空时应该提前判断，否则会触发岩机。

可以支持结构体标签（Struct Tag），方便自定义JSON的键名及忽略某些字段的序列化过程，避免这些字段被序列化到JSON中。

支持缩进且可以自定义缩进字符，将JSON序列化后的内容格式化，方便查看。

默认应该序列化为[]byte字节数组，外部自己转换为字符串。在大部分的使用中，JSON一般以字节数组方式解析、存储、传输，很少以字符串方式解析，因此避免字节数组和字符串的转换，可以提高一些性能。
*/
