package constant_test

import "testing"

// 连续常量的快速赋值
const (
	Monday = 1 + iota // 1 + iota，枚举，每次加1
	Tuesday
	Wednesday
)

// 使用iota，进行左移操作，构造一些强大的枚举常量值生成器
const (
	Readable   = 1 << iota // 0001，1 << iot，每次左移一个二进制位
	Writeable              // 0010
	Executable             // 0100
)

func TestAble(t *testing.T) {
	t.Log(Readable)   // 1，0001
	t.Log(Writeable)  // 2，0010
	t.Log(Executable) // 4，0100
}

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1                                                                              // 0001
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable) // 两数对应位相与，全1为1，有0为0；1为true，0为false
}
