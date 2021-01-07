package array_test

import "testing"

// 数组声明并初始化零值，数组读写
func TestArrayInit(t *testing.T) {
	var arr [3]int             // 数组声明并初始化零值
	arr1 := [4]int{1, 2, 3, 4} // 在数组声明的同时，初始化值
	arr3 := [...]int{1, 3, 4, 5}

	arr1[1] = 5 // 读写数组

	t.Log(arr[1], arr[2]) // 读写数组
	t.Log(arr1, arr3)
}

// 遍历数组元素，3种形式
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5} // 不指定数组元素的个数

	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i]) // 遍历，输出数组的所有元素
	}

	for idx, e := range arr3 {
		t.Log(idx, e) // 使用range，输出数组的所有索引和元素
	}

	for _, e := range arr3 {
		t.Log(e) // 使用_作为无关返回值的占位符
	}
}

/*  使用 for range 遍历数组
for 索引 元素 := range 数组 {
    fmt.Println(索引, 元素)
}
*/

// 数组截取
func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr4[1:2])         // 2
	t.Log(arr4[1:3])         // 2 3
	t.Log(arr4[1:len(arr4)]) // 2 3 4 5
	t.Log(arr4[1:])          // 2 3 4 5
	t.Log(arr4[:3])          // 1 2 3
	// t.Log(arr4[:-1]) // go语言不支持负数类型的索引。invalid slice index -1(index must be non-negative)
	t.Log(arr4[:]) // 全选
}

/*

[a:b]，前闭后开，ab为基数索引。

a[开始索引(包含), 结束索引(不不包含)]

*/
