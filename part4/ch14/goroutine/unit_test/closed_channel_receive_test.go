package goroutine_test

import (
	"fmt"
	"testing"
)

func TestClosedChannelReceive(t *testing.T) {
	ch := make(chan int, 2) // 创建能保存两个整型元素的通道，缓冲大小为2

	// 在通道中放入两个数据，通道装满了
	ch <- 0
	ch <- 1

	// 关闭通道，此时带缓冲通道的数据不会释放，通道也没消失
	close(ch)

	// 遍历所有缓冲的数据，且多遍历一个
	for i := 0; i < cap(ch)+1; i++ { // cap获取通道的容量；这里故意多取一个元素，造成越界访问
		v, ok := <-ch // 从通道取出数据

		fmt.Println(v, ok) // 打印取出的数据状态
	}
}

/*
0 true
1 true
0 false
*/
