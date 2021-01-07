package main

import "fmt"

// 空结构体
type data struct{}

func dummy() *data {
	var c data // 声明c为data结构体类型，基本的实例化

	return &c // 返回函数局部变量的内存地址，是*data类型指针
}

func main() {
	fmt.Println(dummy()) //打印dummy函数的返回值
}

/*
go run -gc flags "-m -l" addrs_escape.go

# command-line-arguments
./addrs_escape.go:11:9: &c escapes to heap
./addrs_escape.go:9:6: moved to heap: c
./addrs_escape.go:15:19: dummy() escapes to heap
./addrs_escape.go:15:13: main ... argument does not escape
&{}
*/
