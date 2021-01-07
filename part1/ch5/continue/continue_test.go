package continue_test

import (
	"fmt"
	"testing"
)

// TestContinueLoop comment
func TestContinueLoop(t *testing.T) {

OuterLoop: // 外层循环的标签
	for i := 0; i < 2; i++ { // 双层循环
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Print(i, j)
				continue OuterLoop // 在continue语句后添加标签，结束当前内层循环，开启下一次外层循环
			}
		}
	}
}

/*
continue语句只能在for循环中使用

continue语句用来结束当前内层循环，并开始下次的外层迭代过程

continue语句后面可以添加标签，表示开始标签对应的循环
*/
