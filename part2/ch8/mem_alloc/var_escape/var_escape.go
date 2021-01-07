package main

import (
	"fmt"
)

// 测试函数入参和返回值
func dummy(a int) int {
	var b int // 声明局部变量（临时变量）
	b = a

	return b // 返回临时变量
}

// 空函数，没有任何入参和返回值
func void() {}

func main() {
	var c int // 测试在main函数中的变量分析
	void()

	fmt.Println(c, dummy(0))
}

/*
./var_escape.go:22:13: c escapes to heap
./var_escape.go:22:22: dummy(0) escapes to heap
./var_escape.go:22:13: main ... argument does not escape
0 0
*/
