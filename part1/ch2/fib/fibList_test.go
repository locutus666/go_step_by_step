package try_test

import "testing"

// var global_var  全局变量/外部变量的声明方式，可以在代码下文或函数内部被赋值或修改

// 产生前10位的Fiblist
func TestFibList(t *testing.T) {
	a := 1 // 自动类型推断的赋值方式，常用于局部变量
	var tmp int
	t.Log(a) // 单元测试中，使用t.Log代替fmt.Println
	for i := 0; i < 9; i++ {
		a, tmp = tmp+a, a // 在一个赋值语句中，可以对多个变量同时赋值，用于交换两个变量，还可以用于运算后交换两个变量
		t.Log(a)
	}
}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
