package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 声明一个Enum类型
type Enum int

const (
	Zero Enum = 0 // Zero是一个Enum类型的常量
)

type cat struct{} // 声明一个空结构体

func TestGetNameKind(t *testing.T) {

	// type cat struct{} // 声明一个空结构体

	typeOfCat := reflect.TypeOf(cat{}) // 获取结构体实例的反射类型对象

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // 获取反射类型对象的名称和种类

	typeOfA := reflect.TypeOf(Zero) // 获取常量Zero的反射类型对象

	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 获取反射类型对象的名称和种类
}

/*
cat struct
Enum int
*/
