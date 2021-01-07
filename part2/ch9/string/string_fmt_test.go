package string_test

import (
	"fmt"
	"testing"
)

var progress int = 2
var target int = 8

func TestStringfmtout(t *testing.T) {
	// 把两个参数格式化输出
	title := fmt.Sprintf("已采集%d个，还需要%d个才能完成任务", progress, target)

	fmt.Println(title)

	pi := 3.14159

	// 正常输出原值
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)

	fmt.Println(variant)
}

/*
已采集2个，还需要8个才能完成任务
月球基地 3.14159 true
*/

func TestStringfmtoutput(t *testing.T) {
	// 声明匿名结构体，并赋初值（结构体实例化）
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}

	fmt.Printf("使用'%%+v', %+v\n", profile)

	fmt.Printf("使用'%%#v', %#v\n", profile)

	fmt.Printf("使用'%%T', %T\n", profile)
}

/*
使用'%+v', &{Name:rat HP:150}
使用'%#v', &struct { Name string; HP int }{Name:"rat", HP:150}
使用'%T', *struct { Name string; HP int }
*/

/*
:=不能用于声明全局变量！只能在函数内部使用，只用来声明临时（局部）变量。

初始化全局变量需使用var关键字，正确操作：var a int = 8

声明结构体时，在{}内的成员变量字段与类型后面都不用加逗号","；在结构体初始化实例赋值时，每个成员变量字段与类型后面都要加","
*/
