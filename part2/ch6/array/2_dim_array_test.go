package array_test

import "testing"

func TestTwoDim(t *testing.T) {
	// 数组，5行2列
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}} // 多维数组声明，并初始化
	var i, j int

	// 输出数组元素
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			t.Logf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
