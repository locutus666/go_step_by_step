package main

import (
	"fmt"
	"sync/atomic"
)

var (
	seq int64
)

func GenID() int64 {
	// atomic.AddInt64(&seq, 1)
	// return seq

	return atomic.AddInt64(&seq, 1) // 尝试原子性地增加序列号，对seq()函数加1
}

func main() {
	for i := 0; i < 10; i++ { // 循环10次，生成10个goroutine调用GenID()函数，同时忽略GenID()的返回值
		go GenID()
	}

	fmt.Println(GenID()) // 单独调用一次GenID()函数
}
