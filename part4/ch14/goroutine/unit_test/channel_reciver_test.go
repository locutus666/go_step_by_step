package goroutine_test

import (
	"fmt"
	"testing"
)

func TestChannekReciver(t *testing.T) {
	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	<-ch

	fmt.Println("all done")
}
