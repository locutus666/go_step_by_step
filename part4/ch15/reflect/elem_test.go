package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestElem(t *testing.T) {
	type cat struct{}

	ins := &cat{} // 实例化cat结构体，得到结构体指针类型*cat

	typeOfCat := reflect.TypeOf(ins) // 使用reflect.TypeOf，获取结构体实例的反射类型对象

	// 使用反射类型对象的.Name()方法和.Kind()方法，显示反射类型对象的类型名和种类
	fmt.Printf("name:'%v' kind:'%v' \n", typeOfCat.Name(), typeOfCat.Kind())

	typeOfCat = typeOfCat.Elem() // 只有新声明变量时才会用到海象运算符，其他情况的赋值使用=。no new variables on left side of :=
	fmt.Printf("element name:'%v' element kind:'%v' \n", typeOfCat.Name(), typeOfCat.Kind())
}

/*
name:'' kind:'ptr'
element name:'cat' element kind:'struct'
*/
