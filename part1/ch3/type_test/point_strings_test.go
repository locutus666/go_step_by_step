package type_test

import "testing"

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a // 使用取址符&，获取变量a的指针，即存放变量a的内存地址
	//aPtr = aPtr + 1  // go不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 使用格式化符%T，获取变量类型
}

func TestStrings(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //string初始化零值是""，不是nil
	t.Log(len(s))
}

/*
1. go不支持任何不同类型的隐式类型转换，也不支持在类型别名条件下，同类型之间的隐式类型转换。go类型转换必须通过显式类型转换

2. go支持指针类型，但不支持指针运算

3. go中的字符串是值类型，默认初始化零值是空字符串""，而不是nil
*/
