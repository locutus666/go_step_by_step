package goroutine_test

import (
	"fmt"
	"testing"
)

func TestBufferChannel(t *testing.T) {
	ch := make(chan int, 3) // 创建有3个整型元素的缓冲通道

	fmt.Println(len(ch)) // 查看当前通道大小

	ch <- 1 // 发送3个整型元素到通道
	ch <- 2
	ch <- 3

	fmt.Println(len(ch)) // 查看当前通道大小
}
