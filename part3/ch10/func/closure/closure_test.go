package func_test

import (
	"fmt"
	"testing"
)

// 准备一个字符串
var str = "hello world"

// 创建一个匿名函数
var foo = func() {
	str = "hello china"
	fmt.Println(str)
}

func TestClosure(t *testing.T) {
	foo()
	fmt.Println(str)
}
