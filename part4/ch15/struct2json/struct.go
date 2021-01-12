package main

import (
	"bytes"
	"reflect"
)

// 在 JSON 格式中，切片是一系列值的序列，以方括号开头和结尾：结构体由键值对组成，以大括号开始和结束 。两种结构的元素均以逗号分隔。
// 将结构体序列化为JSON格式，并输出到缓冲中
func writeStruct(buff *bytes.Buffer, value reflect.Value) error {

	// 取出值的类型对象
	valueType := value.Type()

	// 写入结构体左大括号
	buff.WriteString("{")

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {

		// 获取每个字段的字段值（reflect.Value)
		fieldValue := value.Field(i)

		// 获取每个字段的类型 （reflect.StructField）
		fieldType := valueType.Field(i)

		// 写入字段名左双引号，双引号本身需要使用“＼”进行转义
		buff.WriteString("\"")

		// 写入字段名
		buff.WriteString(fieldType.Name)

		// 写入字段名右双引号和冒号
		buff.WriteString("\":")

		// 递归调用任意值序列化函数writeAny，将fieldValue继续序列化。写入每个字段值
		writeAny(buff, fieldValue)

		// 和切片一样，多个结构体字段间也是以逗号分隔，每个字段尾部写入逗号，最后一个字段不添加
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	// 写入结构体右大括号
	buff.WriteString("}")

	return nil
}
