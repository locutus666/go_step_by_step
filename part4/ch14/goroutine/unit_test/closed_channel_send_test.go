package goroutine_test

import (
	"fmt"
	"testing"
)

func TestClosedChannelSend(t *testing.T) {
	ch := make(chan int)

	close(ch) // 关闭通道

	// 打印通道的指针，容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))

	ch <- 1 // 向关闭的通道发送数据
}

//  panic: send on closed channel，提示panic宕机的原因是给一个关闭了的通道发送数据
