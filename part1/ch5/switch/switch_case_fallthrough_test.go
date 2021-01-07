package controll_test

import (
	"fmt"
	"testing"
)

// 为了兼容一些移植代码，go加入了fallthrough关键字来实现case之后紧接着执行下一个case的功能。
func TestSwitchFallthrough(t *testing.T) {

	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

/*
在Go语言中，case是一个独立的代码块 ，执行完毕后不会像C语言那样紧接着下一个case执行。
为了兼容一些移植代码，go加入了fallthrough关键字来实现case之后紧接着执行下一个case的功能。

新编写的代码不建议使用fallthrough。
*/
