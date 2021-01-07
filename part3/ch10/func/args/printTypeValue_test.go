// 获得可变参数类型 —— 获取每一个参数类型

package func_test

import (
	"bytes"
	"fmt"
	"testing"
)

// 把一系列值传入printTypeValue函数，打印传入参数的值与类型
func printTypeValue(slist ...interface{}) string { // 当可变参数为interface{}类型时，可以传入任何类型的值
	var b bytes.Buffer // 定义字节缓冲变量b，快速连接字符串

	for _, s := range slist { // 遍历slist的每一个元素，slist是interface{}类型
		str := fmt.Sprintf("%v", s) // 使用fmt.Sprintf和"%v"，把interface{}格式的任意值转为字符串

		var typeString string // 声明一个字符串，作为变量的类型名

		switch s.(type) { // switch s.(type)对interface{}类型的值进行类型断言，即判断变量的实际类型
		case bool:
			typeString = "bool" // case多种情况判断，把每种类型对应的字符串赋值到typeString中
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value: ") // 写输出格式

		b.WriteString(str)

		b.WriteString(" type: ")

		b.WriteString(typeString)

		b.WriteString("\n")
	}
	return b.String()
}

func TestGetParamType(t *testing.T) {
	t.Log(printTypeValue(100, "str", true))
}
