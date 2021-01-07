package main

import "fmt"
import "os"

//func main() int {
//	fmt.Println("hello world!")
//    return 1
//}

func main() {
	fmt.Println(os.Args)         // go run hello.go yue，使用os.Args获取命令行参数
	if len(os.Args) > 1 {
		fmt.Println("hello world!", os.Args[1])  // 使用os.Args获取命令行参数。如果它长度大于1，就放在"hello world!"的后边
	}
	fmt.Println("hello world!")
	os.Exit(-1)                  // go通过os.Exit()来返回函数退出码，而不是使用return。0（正常退出，退出码0），-1（异常退出，退出码255）
}
