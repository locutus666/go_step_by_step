package struct_test

import (
	"fmt"
	"testing"
)

// 定义点结构体
type Point struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func TestNoPointReceiver(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{2, 2}

	result := p1.Add(p2)

	fmt.Println(result)
}
