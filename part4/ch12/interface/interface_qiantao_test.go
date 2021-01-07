package interface_test

import (
	"io"
	"testing"
)

// // io包中的写入接口
// type Writer interface {
// 	Write(p []byte)
// }

// // io包中的关闭接口，实现内存资源的释放
// type Closer interface {
// 	Close() error
// }

// // io包中的写入关闭接口，同时拥有Writer和Closer的特性
// type WriteCloser interface {
// 	Writer // 由两个接口嵌套组合，生成扩展接口
// 	Closer
// }

// 声明一个设备结构体，用于模拟一个虚拟设备
type device struct{}

// 向结构体绑定类似io.Writer的Write方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 向结构体绑定类似io.Closer的Close方法
func (d *device) Close() error {
	return nil
}

//
func TestInterfaceQiantao(t *testing.T) {
	// device结构体实例化wc，由于device类型已经实现了io.WriteCloser的所有内嵌接口，因此wc的类型可以看作是io.WriteCloser接口类型（被隐式转换了）
	var wc io.WriteCloser = new(device)

	// wc实例调用Write方法，写入数据
	wc.Write(nil)

	// wc实例调用Close方法，关闭设备
	wc.Close()

	// 声明写入器writeOnly，类型是io.Writer，并赋值device新实例
	var writeOnly io.Writer = new(device)

	// writeOnly实例只调用Write方法，写入数据
	writeOnly.Write(nil)
}
