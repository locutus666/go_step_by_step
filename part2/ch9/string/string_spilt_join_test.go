package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 定义返回变量类型的函数
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 字符串分割与拼接
func TestStringSplit(t *testing.T) {
	s := "/A,B,C"

	parts := strings.Split(s, ",") // 使用逗号，分割字符串，得到一个字符串切片
	t.Log(typeof(parts))
	t.Log(parts)

	for _, part := range parts { // 遍历分割后的字符串，舍弃_
		t.Log(part)
	}

	pinjie := strings.Join(parts, "/") // 使用‘/’，拼接字符串切片。此操作可用于拼接字符串路径
	t.Log(typeof(pinjie))
	t.Log(pinjie)
}

// 字符串与整数的转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10) // 整数到字符串，Itoa：int to a
	t.Log("str" + s)

	// c := 1strconv.Atoi("10") // 从字符串到整数的转换，会返回两个值（转换后的值和err信息）。我们需要对err进行判断，否则直接输出会报错
	// t.Log(c + 10)
	if i, err := strconv.Atoi("10"); err == nil { // 字符串到整数，Atoi：a to int
		t.Log(i + 10)
	}
}
