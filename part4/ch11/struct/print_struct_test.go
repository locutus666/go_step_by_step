package struct_test

import (
	"fmt"
	"testing"
)

// 打印消息类型，传入匿名结构体
func printMsgType(msg *struct {
	id   int // 匿名结构体
	data string
}) { // 函数体
	fmt.Printf("%T\n", msg) // 使用动词%T，打印msg的类型
	fmt.Println(msg)
}

func TestPrintMsgType(t *testing.T) {
	msg := &struct { // 使用&取地址实例化，得*struct指针类型变量
		id   int
		data string
	}{
		1024, // 值列表初始化成员变量
		"hello",
	}

	printMsgType(msg) // 把msg变量传入printMsgType函数
}
