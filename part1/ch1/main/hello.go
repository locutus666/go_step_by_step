package main

import (
	"fmt"
	"os"
)

func main () {
	fmt.Println("hello world!")
	os.Exit(0)                  // go通过os.Exit()来返回函数退出码，而不是使用return。0（正常退出），-1（异常退出）
}


// 一个go语言包中只能有一个package main和一个func main()
// go源文件名不一定要与包名一致