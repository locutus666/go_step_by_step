package struct_test

import (
	"fmt"
	"testing"
)

// 定义结构体Property
type Property struct {
	value int
}

// 定义修改字段值的方法
func (p *Property) SetValue(v int) { //没有返回值，就没有返回类型
	p.value = v
}

// 定义获取字段值的方法
func (p *Property) GetValue() int {
	return p.value
}

func TestPointReceiver(t *testing.T) {
	p := new(Property)        // 实例化Property结构体，得到*T结构体指针类型
	p.SetValue(100)           // 调用Setvalue方法，修改属性值
	fmt.Println(p.GetValue()) // 调用Getvalue方法，获取属性值
}
