package ptr_test

import (
	"fmt"
	"testing"
)

// 使用指针进行数值交换：交换两个指针变量对应的（普通变量）值
func swap(a, b *int) {
	t := *a // 取a指针对应的普通变量值，赋给临时变量t

	*a = *b // 取b指针对应的普通变量值，赋给a指针对应的普通变量

	*b = t // 把a指针对应的普通变量值，赋给b指针指向的变量

}

func TestPtrReviseCode(t *testing.T) {
	x, y := 1, 2

	swap(&x, &y) // 取普通变量的地址，得到指针变量，是指针类型*T，作为swap的入参

	fmt.Println(x, y)
}

// 2 1
