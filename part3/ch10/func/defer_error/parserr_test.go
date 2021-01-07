package func_test

import (
	"fmt"
	"testing"
)

// 声明一个解析错误的结构体ParseError，包含文件名，行号两个成员
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现错误接口error，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line) // 把结构体成员的文件名和行号，格式化字符串返回
}

// 根据文件名和行号，创建一个错误实例
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func TestError(t *testing.T) {
	var e error                     // 声明一个错误接口类型
	e = newParseError("main.go", 1) // 创建一个错误实例，包含文件名和行号

	fmt.Println(e.Error()) // 调用Error方法

	switch detail := e.(type) { // 通过错误断言，取出发生错误的具体类型，获取详细错误信息
	case *ParseError: // 解析错误
		fmt.Printf("Filename:%s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}
