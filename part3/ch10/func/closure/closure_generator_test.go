package func_test

import (
	"fmt"
	"testing"
)

// 创建一个玩家生成器，输入名称，输出生成器
func playGen(name string) func() (string, int) {
	hp := 150 // hp是playGen函数的局部变量（不能被playGen函数的外部访问和修改），也是被匿名函数引用的自由变量。

	return func() (string, int) { // 返回匿名函数
		return name, hp // 匿名函数引用外部变量name和hp，形成闭包
	}
}

func TestGenerator(t *testing.T) {
	generator := playGen("high noon") // 生成一个玩家生成器实例（闭包）
	name, hp := generator()           // 调用闭包，获取记忆变量

	fmt.Println(name, hp)
}
