package func_test

import (
	"fmt"
	"testing"
)

// 定义累加器生成函数Accumulate
func Accumulate(value int) func() int {
	return func() int { // 返回一个闭包，这个闭包是一个匿名函数引用了装饰者函数的入参
		value++
		return value // 返回被修改后的自由变量
	}
}

func TestClosureMemory(t *testing.T) {
	accumulator := Accumulate(1) // 创建一个闭包实例accumulator，引用入参1
	fmt.Println(accumulator())   // 调用闭包实例，并打印闭包中被修改后的变量，证明闭包的记忆性
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator) // 打印闭包的内存地址

	accumulator2 := Accumulate(10)   // 创建一个闭包实例accumulator，引用入参10
	fmt.Println(accumulator2())      // 调用闭包实例，并打印闭包中被修改后的变量，证明闭包的记忆性
	fmt.Printf("%p\n", accumulator2) // 打印闭包的内存地址
}
