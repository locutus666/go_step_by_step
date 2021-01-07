package interface_test

import (
	"fmt"
	"testing"
)

func printType(v interface{}) {
	switch v.(type) { // v.(type)是类型分支的典型写法，代码经过switch时，会判断空接口v中包含的具体类型，从而进行类型分支跳转
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func TestTypeSwitch(t *testing.T) {
	printType(1024)
	printType("pig")
	printType(true)
}
