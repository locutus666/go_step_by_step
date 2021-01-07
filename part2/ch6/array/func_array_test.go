package array_test

import (
	"fmt"
	"testing"
)

// 一维数组传参
func TestFuncArray(t *testing.T) {
	var balance = []int{1000, 2, 3, 17, 50} // 切片长度为5
	var avg float32

	// 切片作为参数，传递给函数
	avg = getAverage(balance, 5)

	// 格式化输出返回的平均值
	t.Logf("平均值为: %f ", avg) //
}

func getAverage(arr []int, size int) float32 { // 输出是float32类型
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size) // 显式类型转换
	return avg                         // 返回变量
}

// 二维数组传参
func prt(arr [][]float32, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(arr[i][0])
	}
}

func TestMain(t *testing.T) {
	var arr = [][]float32{{-1, -2}, {-3, -4}, {-5}}
	prt(arr, 3)
}

// 形参设定数组长度
func TestMainSet(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5}
	setArray(array) // 未定义长度的数组，只能传给不限制数组长度的函数

	var array2 = [5]int{1, 2, 3, 4, 5} // 定义了长度的数组，只能传给限制了相同数组长度的函数
	setArray2(array2)
}

func setArray(params []int) {
	fmt.Printf("params array length of setArray is : %d", len(params)) // fmt.Printf格式化输出
}

func setArray2(params [5]int) {
	fmt.Printf("params array length of setArray2 is : %d", len(params))
}
