package interface_test

import (
	"fmt"
	"testing"
)

// 定义一个数据写入器接口，这个接口只有一个方法WriteData
type DataWriter interface {
	WriteData(data interface{}) error // 输入interface{}类型的变量data，返回一个error表示是否出错
}

// 定义文件结构体，用于实现DataWriter接口
type file struct{}

func (d *file) WriteData(data interface{}) error { // 绑定结构体方法，实现DataWriter接口的WriteData方法。输入interface{}类型的变量data，返回一个error表示是否出错
	fmt.Println("WriteData:", data)
	return nil
}

func TestDataWriter(t *testing.T) {

	f := new(file) // 实例化file结构体，类型为*file

	var writer DataWriter // 声明一个DataWriter接口类型的变量write

	writer = f // 把f（*file结构体指针类型）赋值给writer

	writer.WriteData("data") // 使用DataWriter接口进行数据写入（调用WriteData方法）
}

//WriteData: data
