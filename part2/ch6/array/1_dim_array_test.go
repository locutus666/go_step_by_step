package array_test

import "testing"

func TestArray(t *testing.T) {
	var n [10]int // n是一个长度为10的数组
	var i, j int

	// 初始化数组n
	for i = 0; i < 10; i++ {
		t.Logf("i = %d\n", i)
		n[i] = i + 100 // 设置元素为 i + 100
	}

	// 输出数组n元素
	for j = 0; j < 10; j++ {
		t.Logf("Element[%d] = %d\n", j, n[j]) // 使用%d，格式化输出整数；/n表示换行
	}
}
