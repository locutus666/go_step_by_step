package main

import (
	"bytes"
	"reflect"
)

// 切片序列化
// writeAny函数中会调用writeSlice函数，把切片类型转换为JSON格式的字符串，并将数据写入缓冲中
func writeSlice(buff *bytes.Buffer, value reflect.Value) error {

	// 写入切片开始标识 "["和结束标识 "]" 。
	buff.WriteString("[")

	// 遍历每个切片元素
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		// 写入每个切片元素
		writeAny(buff, sliceValue)

		// 每个元素尾部写入逗号，最后一个字段不添加
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("]")

	return nil
}

// JSON格式规定：每个数组成员由逗号分隔且最后一个元素后不加逗号。
