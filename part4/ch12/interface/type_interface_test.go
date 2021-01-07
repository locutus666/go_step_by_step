package interface_test

import (
	"io"
	"testing"
)

// // 接口定义方
// // io包中的写入接口
// type Writer interface {
// 	Write(p []byte)
// }

// // io包中的关闭接口，实现内存资源的释放
// type Closer interface {
// 	Close() error
// }

// 接口实现方：使用Socket实现io.Writer接口
type Socket struct{}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

// 接口使用方：在代码中使用上述接口
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

func usingCloser(closer io.Closer) {
	closer.Close()
}

func TestTypeInterface(t *testing.T) {
	s := new(Socket) // 实例化Socket结构体

	usingWriter(s)
	usingCloser(s)
}
