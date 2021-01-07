package struct_test

import (
	"testing"
)

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func TestMemberChongtu(t *testing.T) {
	c := &C{} // 内嵌结构体实例化
	c.A.a = 1 // 成员变量类型已经定义，所以不用:=类型推断赋值，而用=
	t.Log(c)

	// c.a =1
	// t.Log(c)
}
