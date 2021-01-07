package if_test

import (
	"fmt"
	"testing"
)

func TestIfIf(t *testing.T) {
	// 定义局部变量
	var a int = 100
	var b int = 200

	// 条件判断
	if a == 100 {
		// if条件语句为true，则执行
		if b == 200 {
			// if条件语句为true，则执行
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}

	fmt.Printf("a 值为 %d\n", a)
	fmt.Printf("b 值为 %d\n", b)
}

func TestIfIfElse(t *testing.T) {
	// var count,c int   // 定义了变量不使用，也会报错
	var count int
	var flag bool
	count = 1

	// while(count<100) {    // go没有while
	for count < 100 {
		count++
		flag = true
		//注意tmp变量  :=
		for tmp := 2; tmp < count; tmp++ {
			if count%tmp == 0 {
				flag = false
			}
		}

		// 每一个 if else 都需要加入括号，同时 else 位置不能在新一行
		if flag == true {
			fmt.Println(count, "素数")
		} else {
			continue
		}
	}
}

func TestIfIfElseElse(t *testing.T) {
	var a int
	var b int

	fmt.Printf("请输入密码： \n")
	fmt.Scan(&a)

	if a == 5211314 {
		fmt.Printf("请再次输入密码：")
		fmt.Scan(&b)
		if b == 5211314 {
			fmt.Printf("密码正确，门锁已打开")
		} else {
			fmt.Printf("非法入侵，已自动报警")
		}
	} else {
		fmt.Printf("非法入侵，已自动报警")
	}
}
