// 遍历可变参数列表 —— 获取每一个参数的值

package func_test

import (
	"bytes"
	"testing"
)

// 定义函数joinStrings，可变参数数量为0～n，类型为字符串
func joinStrings(slist ...string) string { //入参slist是可变参数，类型为[]string切片，即此函数只接受字符串作为参数
	var b bytes.Buffer // 定义字节缓冲变量b，快速连接字符串

	for _, s := range slist { // 遍历可变参数slist，类型为[]string切片
		b.WriteString(s) // 把遍历出的字符串值，连续写入字节变量b
	}

	return b.String() // 把字符缓冲变量中的数据转换为字符串
}

// 输入多个字符串，把他们连成一个字符串
func TestParamter(t *testing.T) {
	t.Log(joinStrings("Pig ", "and", " rat"))
	t.Log(joinStrings("hammer", " mom", " and", " hawk"))
}
