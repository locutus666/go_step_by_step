package controll_test

import (
	"fmt"
	"testing"
)

// case后面跟常量
func TestSwitchCaseConst(t *testing.T) {
	var a = "hello"

	switch a {
	// 当i的类型是以下各种情况时
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	case "mum", "daddy": // 一分支多值，不同的值或表达式使用逗号分隔
		fmt.Println(3)
	default:
		fmt.Println(0)
	}
}

// case后面跟变量
func TestSwitchCaseVar(t *testing.T) {
	var x interface{}

	// switch之后可以有变量初始化语句
	switch i := x.(type) {
	// 当i的类型是以下各种情况时
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

// case后面跟表达式
func TestSwitchCaseExpression(t *testing.T) {

	var grade string // 定义局部变量
	var marks int

	fmt.Println("输入一个数字：")
	fmt.Scan(&marks) // 从命令行获取一个数字，赋给变量marks

	switch {
	case marks >= 90:
		grade = "A"
	case marks < 90 && marks >= 80:
		grade = "B"
	case marks < 80 && marks >= 70:
		grade = "C"
	case marks < 70 && marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("中等\n")
	case grade == "D":
		fmt.Printf("及格\n")
	default:
		fmt.Printf("不及格\n")
	}

	fmt.Printf("你的等级是 %s\n", grade)
}

/*
分支选择可以理解为一种批量的if语旬。使用switch语句可以方便地判断大量的数值常量/变量，字符串或表达式。

Go语言规定，每个switch语句可以有多个case分支，但只能有一个default分支。

在Go语言中，case是一个独立的代码块，执行完毕后不会像C语言那样紧接着下一个case执行。

switch语句中的每一个case与case之间是独立的代码块，不需要通过break语句跳出当前case代码块，否则，代码会继续执行到下一行。

语言的运行效率并不能直接决定最终的效率，I/O效率现在是最主要的问题。因此， Go语言中的switch语法设计尽量以使用方便为主。
*/
