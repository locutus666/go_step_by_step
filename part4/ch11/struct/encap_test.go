package struct_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string // field type
	Name string // 字段 类型
	Age  int    // 成员 类型，这三种叫法一致
}

// // 把String()方法附加到Employee结构体上
// func (e Employee) PString() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) // 输出入参的内存地址
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 把PString()方法附加在Employee的指针类型上
func (e *Employee) PString() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	f := Employee{"0", "Bob", 20} // 结构体实例化和初始化

	f1 := Employee{Name: "Mike", Age: 30}

	f2 := new(Employee) // 实例化，返回结构体指针
	f2.Id = "2"         // 初始化
	f2.Age = 22
	f2.Name = "Rose"

	t.Log(f)     // {0 Bob 20}
	t.Log(f1)    // { Mike 30}
	t.Log(f1.Id) // f1的ID为空

	t.Log(f2) // &{2 Rose 22}

	t.Logf("f is %T", f)   // t.Logf("%T", v)表示取变量类型，f是Employee类型
	t.Logf("f is %T", &f)  // &f是Employee的指针类型
	t.Logf("f2 is %T", f2) // f2是Employee的指针类型
}

func TestStructOperations(t *testing.T) {
	g := Employee{"0", "Bob", 20} // 通过类型实例与类型指针访问，结果一样
	// g := &Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&g.Name)) // 输出函数调用时g的内存地址
	//t.Log(f.String())
	t.Log(g.PString()) // 调用结构体方法，返回变量的值和内存地址。两次返回的内存地址一样，证明实例成员变量没有进行值复制；两次返回的内存地址不一致，证明实例成员变量进行了值复制
}

// 指针类型接收器（常用），不进行成员变量值的内存复制；一般类型接收器，成员变量值会进行内存复制，不建议使用
