package type_test

import (
	"fmt"
	"testing"
)

// 将NewInt定义为int类型
type NewInt int

// 把int取一个别名IntAlias
type IntAlias = int

func TestType(t *testing.T) {
	// 把a声明为NewInt类型
	var a NewInt

	// 查看a的类型名
	fmt.Printf("a type: %T\n", a) // a type: type_test.NewInt

	// 将a2声明为IntAlias类型
	var a2 IntAlias

	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2) // a2 type: int
}
