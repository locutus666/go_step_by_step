package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 函数闭包：先在函数体外部声明自由变量，然后在函数体内部引用函数体外部的自由变量，进行修改，能改变这个自由变量的值。

// type cat struct {
// 	Name string
// 	Type int `json:"type" id:"110"` // 带有结构体tag的成员变量字段
// }

// ins := cat{
// 	Name: "mimi",
// 	Type: 1
// }

// typeOfCat := reflect.TypeOf(ins)  // 获取结构体实例的反射类型对象

func TestTypeMember(t *testing.T) {

	// 声明带有两个成员的cat结构体
	type cat struct {
		Name string
		Type int `json:"type" id:"110"` // 带有结构体标签（struct tag）的成员变量字段Type。这个成员后面带有``字符串，一般用于给字段添加自定义信息，方便其他模块的处理
	}

	ins := cat{
		Name: "mimi",
		Type: 1, // 结构体实例化时，成员字段的每个末尾都要加, 否则报错missing ',' before newline in composite literal。结构体标签属于类型信息，不需要赋值。
	}

	typeOfCat := reflect.TypeOf(ins) // 获取结构体实例ins的反射类型对象

	//
	for i := 0; i < typeOfCat.NumField(); i++ { // 使用反射类型对象的NumField方法，返回反射类型对象对应的结构体成员变量字段数量。i为结构体成员变量的索引，当i小于反射类型对象的结构体成员变量字段数量时，i++

		fieldType := typeOfCat.Field(i)                                  // 根据索引，返回索引对应反射类型对象成员变量的结构体structField
		fmt.Printf("name:%v tag:'%v' \n", fieldType.Name, fieldType.Tag) // 访问返回的结构体structField成员变量Name和Tag
	}

	if catType, ok := typeOfCat.FieldByName("Type"); ok { // 根据给定字符串Type，返回字符串对应的结构体字段信息
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id")) // 使用structField中Tag的Get方法，获取结构体标签的字段
	}
}

/*
name:Name tag:''
name:Type tag:'json:"type" id:"110"'
type 110
*/
