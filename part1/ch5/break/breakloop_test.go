package break_test

import (
	"fmt"
	"testing"
)

// TestBreakLoop comment
func TestBreakLoop(t *testing.T) {

OuterLoop: // 外层循环的标签
	for i := 0; i < 2; i++ { // 双层循环
		for j := 0; j < 5; j++ {
			switch j { // 使用switch进行数值分支判断
			case 2:
				fmt.Println(i, j)
				break OuterLoop // 退出OuterLoop对应的循环，即结束双层循环
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}

/*
break语句用来结束for、 switch和select的代码块

break语句后面可以添加标签，表示退出标签对应的代码块

标签必须定义在对应的for、switch和select代码块上
*/
