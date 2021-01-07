package struct_test

import "testing"

// 自定义类型MyInt
type MyInt int

// 为Myint添加IsZero方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为Myint添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func TestAddMethodToType(t *testing.T) {
	var b MyInt
	t.Log(b.IsZero())

	b = 1
	t.Log(b.Add(2))
}
