package type_test

import "testing"

// type alias
type MyInt = int
type MyInt1 = MyInt

// type definition
type MyInt2 int64

// 采用传统type definition定义的新类型在其变量被赋值时，需对右侧变量进行显式类型转换，否则编译器就会报错。
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	//b = a           // 隐式类型转换，cannot use a (type int) as type int64 in assignment
	b = int64(a) // 显式类型转换

	var c MyInt2
	//c = b // go语言的类型转换很严格，即使类型定义的"隐式类型转换"也不允许，强调"显式类型转换"。
	c = MyInt2(b)
	t.Log(a, b, c)
}

// type alias并未创造新类型，只是源类型的“别名”，在类型信息上与源类型一致，因此可以直接赋值
func TestTypeAlias(t *testing.T) {
	var i int = 5
	var mi MyInt = 6
	var mi1 MyInt1 = 7

	mi = i   // ok
	mi1 = i  // ok
	mi1 = mi // ok

	t.Log(i, mi, mi1)

}
