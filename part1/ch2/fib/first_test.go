package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}

// The master has failed more times than the beginner has tried.

/*
单元测试：go源文件以 _test 结尾：xxx_test.go

package xxx_test

func TestXxx(t *testing.T) {
	...

	t.Log()   // t.Log输出，t.Logf()格式化输出

}

其中t是变量名，T是testing包中的一个类型，*testing.T是testing包中T类型的指针类型

在命令行运行单元测试：go test -v xxx_test.go
*/
