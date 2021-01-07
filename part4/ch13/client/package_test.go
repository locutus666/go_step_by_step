package client_test

import (
	"testing"

	"github.com/locutus666/go_step_by_step/part4/ch13/series" // 导入外部包，包括GOPATH路径下的，vendor目录下的，GOROOT路径下的
)

// 使用外部包的函数
func TestOuterPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.Square(5))

	// t.Log(series.square(5)) // 访问包的私有函数，会报错cannot refer to unexported name series.square，undefined: series.square
}
