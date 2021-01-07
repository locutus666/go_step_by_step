package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

var (
	count      int        // 某个逻辑中使用的变量，如包级别变量、结构体成员字段变量
	countGuard sync.Mutex // 变量名+Guard，以表示这个互斥锁用于保护这个变量。一般情况下，把互斥锁的粒度设置得越小越好，降低共享访问的等待时间。
)

// 获取一个count值，通过这个函数可以并发安全地访问变量count
func GetCount() int {
	countGuard.Lock() // 对互斥量countGuard加锁。一旦互斥量被加锁，如果另一个goroutine尝试继续加锁时会发生阻塞，直到互斥量解锁

	defer countGuard.Unlock() // 在函数退出时解锁

	return count
}

// 设置一个count值，同样使用加锁解锁，保证修改count值的过程是一个原子过程，不会发生并发访问冲突
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func TestSyncMutex(t *testing.T) {
	for i := 0; i < 10; i++ {
		SetCount(i)             // 进行并发安全设置一个count值
		fmt.Println(GetCount()) // 进行并发安全获取
	}
}
