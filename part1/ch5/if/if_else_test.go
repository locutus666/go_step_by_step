package if_test

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	var s int // 声明变量s是需要判断的数
	fmt.Println("输入一个数字：")
	fmt.Scan(&s)

	if s%2 == 0 { // s 除以 2 的余数是否等于 0
		fmt.Print("s是偶数\n") // 如果成立
	} else { // 否则
		fmt.Print("s不是偶数\n")
	}

	fmt.Print("s 的值是：", s)
}
