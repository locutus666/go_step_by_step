package series

import "fmt"

// 在main函数被调用前，所有依赖包的init函数都会被依次执行
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 公有函数
func Square(n int) int {
	return n * n
}

// 包私有函数
func sum(a int, b int) int {
	return a + b
}

// 公有函数
func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret
}
