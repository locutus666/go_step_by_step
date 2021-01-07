package loop_test

import (
	"testing"
)

// for循环的形式
func TestFor(t *testing.T) {
	var b int = 15
	var a int // 初始零值

	// 第一种形式，标准C语言形式的循环
	for a := 0; a < 10; a++ {
		t.Logf("a 的值为: %d\n", a)
	}

	// 第二种形式，省略初始语句的循环
	step := 2
	for ; step > 0; step-- {
		t.Log(step)
	}

	// 第三种形式，只有结束语句的循环，死循环，靠break跳出
	var i int
	for ; ; i++ { // 没有设置i的初始值
		if i > 10 {
			break
		}

		t.Log("程序完成")
	}

	// 第四种形式，只有循环条件的循环，死循环，靠break跳出
	// 类似于其他编程语言中的while，在while后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环
	for a < b {
		a++
		t.Logf("a 的值为: %d\n", a)
	}

	// 第五种形式，无限循环，靠break跳出
	// 在收发处理中，无限循环较为常见，但需要无限循环有可控的退出方式来结束循环。
	for { // 忽略for的所有语句，此时for执行无限循环
		if a > 10 {
			break
		}

		a++ // 把i++从for的结束语句放置到函数体末尾是等效的，这样编写的代码更具有可读性。
	}
}

/*
for 初始语句；条件表达式；结束语句｛
	循环体代码
}

循环体不停地进行循环，直到条件表达式返回false时自动退出循环，执行for的"}"之后的语句。
*/
