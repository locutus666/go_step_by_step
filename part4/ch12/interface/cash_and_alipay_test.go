package interface_test

import (
	"fmt"
	"testing"
)

// 接口实现
type Alipay struct{} // 电子支付方式

func (a *Alipay) CanUseFaceID() {} // 为Alipay添加CanUseFaceID方法，表示电子支付支持刷脸

type Cash struct{} // 现金支付方式

func (a *Cash) Stolen() {} // 为Cash类型添加Stolen方法，表示现金支付方式会出现偷窃现象

// 接口定义
type ContainCanUseFaceID interface { // 具备刷脸特性的接口
	CanUseFaceID()
}

type ContainStolen interface { // 具备被偷特性的接口
	Stolen()
}

// 接口调用方
func print(payMethod interface{}) { // 传入支付方式的接口
	switch payMethod.(type) { // 判断一种支付方法具备哪些特性
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func TestCashAndAlipay(t *testing.T) {
	print(new(Alipay))
	print(new(Cash))
}

/*
*interface_test.Alipay can use faceid
*interface_test.Cash may be stolen
 */
