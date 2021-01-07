package func_test

import (
	"os"
	"testing"
)

// 根据文件名，获取某个文件的大小
func fileSize(filename string) int64 {
	f, err := os.Open(filename) // 根据文件名打开文件，返回文件句柄f和err

	if err != nil { // 如果打开文件时出错，返回文件大小为0
		return 0
	}

	info, err := f.Stat() // 获取文件状态信息和err

	if err != nil { // 如果获取信息时出错，关闭文件并返回文件大小为0
		f.Close()
		return 0
	}

	size := info.Size() // 获取文件大小
	f.Close()           // 关闭文件

	return size // 返回文件大小
}

// 使用defer延迟释放文件句柄
func filesize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil { // 一旦文件打开错误，f将为空
		return 0
	}

	defer f.Close() // 延迟调用Close，在下面代码执行完之后，最后关闭文件

	info, err := f.Stat()

	if err != nil {
		return 0
	}

	size := info.Size()

	return size
}

func TestFileSize(t *testing.T) {
	filesize("file.txt") // 在Go语言中不倾向于使用单引号来表示字符串，请根据需要使用双引号或反引号。否则会报错illegal rune literal
}

// 文件操作需要经过打开文件，获取和操作文件信息、关闭文件等步骤。如果打开后而不关闭文件，进程会一直存在而不退出，继而占用系统资源
// 每一步操作都需要错误处理，而每一步处理都可能会造成退出，因此需要在程序退出时释放文件资源。而我们需要密切关注，在程序可能的退出处，正确关闭文件

// Go语言的单引号一般用来表示rune literal，即码点字面量(应该是Unicode编码表)。
