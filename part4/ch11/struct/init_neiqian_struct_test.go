package struct_test

import (
	"testing"
)

// 车轮，基础构件
type Wheel struct {
	Size int
}

// 引擎，基础构件
type Engine struct {
	Power int
	Type  string
}

// 车，整体结构，结构体内嵌
type Car struct {
	Wheel
	Engine
}

func TestInitNeiqianStruct(t *testing.T) {
	c := Car{
		Wheel: Wheel{
			Size: 18,
		},
		Engine: Engine{
			Power: 143,
			Type:  "1.4T",
		},
	}
	t.Logf("%+v\n", c)
}
