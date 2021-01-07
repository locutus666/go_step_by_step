package ptr_test

import (
	"fmt"
	"testing"
)

func TestGetPtrValue(t *testing.T) {
	// var house = "mabuli point 10880, 90265"
	// house := "mabuli point 10880, 90265"
	var house string = "mabuli point 10880, 90265" // 初始化一个字符串类型的变量house

	ptr := &house // 对字符串变量取内存地址，得到字符串指针类型，赋给指针变量ptr，是*string类型

	fmt.Printf("ptr type: %T\n", ptr) // 获取指针变量的类型
	fmt.Printf("address: %p\n", ptr)  // 获取指针变量的内存地址

	value := *ptr // 对指针变量取值

	fmt.Printf("value type: %T\n", value)
	fmt.Printf("address: %s\n", value)
}

/*
ptr type: *string
address: 0xc0000ae030
value type: string
address: mabuli point 10880, 90265
*/
