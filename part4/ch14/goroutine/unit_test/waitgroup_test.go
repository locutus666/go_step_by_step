package goroutine_test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup // 声明一个等待组。一组等待任务只需要一个等待组，而不是每个任务都需要一个等待组。

	urls := []string{ // 准备一个网站的字符串切片
		"http://www.github.com/",
		"https://www.github.com/",
		"https://www.golangtc.com",
	}

	// 遍历字符串切片中的地址
	for _, url := range urls {

		wg.Add(1) // 每个并发任务开始时，把等待组加1

		// 启动一个匿名函数的并发任务
		go func(url string) {
			defer wg.Done() // 每个并发的匿名函数结束时，都会执行这一句，表示任务完成。等效于wg.Add(-1)

			_, err := http.Get(url) // 使用Get函数访问url，这个操作会一直阻塞（持续访问），直到网站响应或超时
			fmt.Println(url, err)   // 在网站响应或超时后，打印这个网站的地址或错误

		}(url) // 通过参数url入参，为了避免url变量通过闭包引用后被修改的问题

		wg.Wait() // 等待所有任务完成后，wg.Wait()就会停止阻塞。

		fmt.Println("over")
	}
}

/*
http://www.github.com/ <nil>
over

https://www.github.com/ <nil>
over

https://www.golangtc.com <nil>
over
*/
