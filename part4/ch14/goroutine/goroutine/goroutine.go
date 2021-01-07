package main

import (
	"fmt"
	"runtime"
	"time"
)

func running() {
	var times int
	for { // 构建一个无限循环
		times++
		fmt.Println("tick", times) // 输出times变量的值

		time.Sleep(time.Second) // time.Sleep延时1秒，之后继续循环
	}
}

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	go running() // 使running()函数和main()函数并发执行

	var input string
	fmt.Scanln(&input) // 接收用户输入，直到按回车键把命令行输入内容写入input变量
	fmt.Println(input) // 输出input内容
}

/*
在go语言中，runtime运行时会默认为main()函数创建一个goroutine。

在本例中，main()函数的goroutine创建运行后，执行到go running()时，go调度系统又为running()函数
函数创建了一个goroutine，running()函数在自己的goroutine中运行。

此时通过go语言的调度机制，程序有两个goroutine同时运行。
*/
