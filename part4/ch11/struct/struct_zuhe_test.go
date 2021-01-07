package struct_test

import (
	"fmt"
	"testing"
)

// 天上飞的
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 地上走的
type Walking struct{}

// 给Walking结构体添加Walk方法
func (w *Walking) Walk() {
	fmt.Println("can walk")
}

// 结构体内嵌
// 人类
type Human struct {
	Walking
}

// 鸟类
type Bird struct {
	Walking
	Flying
}

func TestStructZuhe(t *testing.T) {
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}
