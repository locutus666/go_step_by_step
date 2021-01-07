package if_test

import (
	"testing"
)

func TestIf(t *testing.T) {
	// 定义局部变量
	var a int = 10

	// 使用if语句，判断条件是布尔表达式
	if a < 20 {
		// 如果条件为 true 则执行以下语句
		t.Logf("a 小于 20\n")
	}

	t.Logf("a 的值为 %d\n", a)
}
