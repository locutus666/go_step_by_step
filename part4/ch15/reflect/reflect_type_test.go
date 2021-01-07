package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType(t *testing.T) {
	var a int

	typeOfA := reflect.TypeOf(a)                // 取得变量a的反射类型对象，是reflect.Type类型
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 获取反射typeOfA的类型名（使用.Name()函数）和种类（使用.Kind()函数）
}

// int int
