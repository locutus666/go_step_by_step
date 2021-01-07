package type_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func TestStructAlias(t *testing.T) {
	// 声明变量a为车辆类型，即把Vechicle实例化为a
	var a Vehicle

	// 显式调用结构体FakeBrand的Show方法
	// 在调用Show()方法时，证明FakeBrand的本质确实是Brand类型。如果两个类型都有Show()方法，会有歧义。
	// 因此只有Brand类型有Show()方法，FakeBrand的本质是Brand类型。
	a.FakeBrand.Show()
	a.Brand.Show()

	// 使用反射，取变量a的类型反射对象，以查看其成员类型
	ta := reflect.TypeOf(a)

	// 遍历a的所有结构体成员变量
	for i := 0; i < ta.NumField(); i++ {

		// a的成员信息
		f := ta.Field(i)

		// 打印Vehicle结构体类型所有成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}

/*
FakeBrand是Brand的一个类型别名。在Vehicle类型中入FakeBrand和Brand，并不意味着嵌入两个Brand。
FakeBrand类型会以名字的方式保留在Vehicle的成员中。

如果尝试把a.FakeBrand.Show()改为a.Show()，编译器将发生报错：ambiguous selector a.Show
在调用Show()方法时，证明FakeBrand的本质确实是Brand类型。因为两个类型都有Show()方法，会发生歧义。
*/
