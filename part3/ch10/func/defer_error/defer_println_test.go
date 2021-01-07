package func_test

import (
	"fmt"
	"testing"
)

// 第一个被defer的语句，最后执行；最后被defer的语句，第一个被执行
func TestDefer(t *testing.T) {
	fmt.Println("defer begin")

	defer fmt.Println(1) // 最后被执行

	defer fmt.Println(2) // 第二个被执行

	defer fmt.Println(3) // 首先被执行

	fmt.Println("defer end")
}

/*
defer语句会延迟处理之后紧随的go语句。

代码的defer延迟顺序与最终执行顺序，是相反的的。

defer延迟调用发生在对应函数结束时。函数结束可以是正常返回，也可以是发生panic宕机
*/
