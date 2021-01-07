package constant_test

import (
	"fmt"
	"testing"
)

// 使用类型定义，声明芯片类型；类型定义与类型别名只差一个“=”
type ChipType int

// 声明None为ChipType类型，值为0
// 声明CPU为ChipType类型，值为1
// 声明GPU为ChipType类型，值为2
const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

// 定义返回字符串的方法String()，把枚举值转换为字符串
// 非指针类型接收器，c为接收器变量，ChipType为接收器类型
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func TestConstant(t *testing.T) {
	// 以整形方式显示，输出CPU值
	fmt.Printf("%s %d", CPU, CPU)
}
