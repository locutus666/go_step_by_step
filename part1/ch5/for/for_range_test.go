package loop_test

import (
	"testing"
)

// 遍历数组、切片 一一 获得索引和元素
func TestForRangeSlice(t *testing.T) {
	// 没有元素的序号位要补0
	numbers := [6]int{1, 2, 3, 5}
	for key, value := range numbers {
		t.Logf("第 %d 位 x 的值 = %d\n", key, value)
	}

	for key, value := range []int{1, 2, 3, 4} { // key和value分别代表切片的下标和下标对应的值。
		t.Logf("key: %d, value: %d\n", key, value)
	}
}

// 遍历字符串 一一 获得字符
func TestForRangeString(t *testing.T) {
	str := "hello, 你好!"
	for key, value := range str { // key和value分别代表字符串的索引（baseO）和字符串中的每一个字符。
		t.Logf("key: %d, value:0x%x\n", key, value) // 变量value，类型是rune，实际上就是int32 ，以十六进制打印出来就是字符的编码。
	}
}

// 遍历map 一一 获得map的键和值
func TestForRangeMap(t *testing.T) {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for key, value := range m { // 对map遍历时，输出键值是无序的。如果需要有序的键值对输出，需要对结果进行排序。
		t.Log(key, value)
	}
}

// 遍历通道（channel） 一一 接收通道数据
// 在遍历通道时，一次只输出一个值，即管道内类型对应的数据。
func TestForRangeChannel(t *testing.T) {
	c := make(chan int)

	go func() { // 声明一个goroutine，往通道中推送数据1、2、3，然后结束并关闭通道。
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}() // 在声明结束后，goroutine马上被并行执行。

	for v := range c { // 使用for range对通道c进行遍历，其实就是不断地从通道中取数据，直到通道被关闭。
		t.Log(v)
	}
}

// 在遍历中选择希望获得的变量 一一 匿名变量（下画线）
// 匿名变量可以理解为一种占位符，它不会进行内存空间分配，也不会占用一个变量的名字。
func TestForRangeMap_(t *testing.T) {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for _, value := range m { // 对map遍历时，输出键值是无序的。如果需要有序的键值对输出，需要对结果进行排序。
		t.Log(value)
	}

	for key := range []int{1, 2, 3, 4} { // value是空变量，没有被使用，移除了_
		t.Logf("key: %d\n", key)
	}
}

/*
for包含初始化语句、条件表达式、结束语旬，这3个部分均可缺省。

for range支持对arry、slice、string、map、channel进行遍历操作。

在需要时，可以使用匿名变量对for range的变量进行选择。
*/
