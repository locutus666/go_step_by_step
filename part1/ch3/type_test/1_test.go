package type_test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

// int占用字节测试
func TestCpu(t *testing.T) {
	cpu := runtime.GOARCH
	int_size := strconv.IntSize

	fmt.Println(cpu, int_size)
	//t.Log(cpu, int_size)
}

// string初始化测试
func TestString(t *testing.T) {
	var s string
	t.Logf("%T %q", s, s)
}
