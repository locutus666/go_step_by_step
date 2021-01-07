package func_test

import (
	"fmt"
	"testing"
)

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	for _, a := range rawList { // 遍历可变参数rawList，raw原始的
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	rawPrint(slist...) // 调用实际打印函数，传参slist，类型为[]interface{}
}

func TestVariaDictTransfer(t *testing.T) {
	print(1, 2, 3) // 调用打印函数，传入实参
}

// 可变参数使用...进行传递，与切片使用append连接是同一个特性
