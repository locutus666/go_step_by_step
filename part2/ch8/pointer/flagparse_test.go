package ptr_test

import (
	"flag"
	"fmt"
	"testing"
)

// 使用flag包，把命令行参数解析传入go代码中提前定义好的一些变量
var mode = flag.String("mode", "", "process mode") // 使用flag.String，定义一个mode变量，类型是*string

func TestFlagParse(t *testing.T) {
	flag.Parse() // 解析命令行参数，并把命令行参数值解析入上述变量mode中

	fmt.Println(*mode) // mode是指针变量，使用*(*T)从指针变量中获取（指针地址指向的）原普通变量的值
}

// go test -v flagparse_test.go --mode=fast

/*
使用flag包定义命令行参数名，返回的是指针变量，格式：flag.String("命令行参数名称", "参数默认值", "参数说明")
*/
