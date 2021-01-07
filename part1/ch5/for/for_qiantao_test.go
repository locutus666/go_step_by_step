package loop_test

import (
	"fmt"
	"testing"
)

// 九九乘法表
func TestForQianTao(t *testing.T) {
	for i := 1; i <= 9; i++ { // 遍历行，决定处理第几行。i控制行，以及计算的最大值
		for j := 1; j <= i; j++ { // 遍历列，决定处理对应行的哪些列。j控制每行的计算列数
			t.Logf("%d*%d=%d ", j, i, j*i) // 每个计算式后边空出一格，便于罗列计算式
		}

		fmt.Println("") // 每执行完一行的所有列计算，就换行
	}
}
