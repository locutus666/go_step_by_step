package func_test

import (
	"sync"
	"testing"
)

var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex // map不是并发安全的，因此需要准备一个sync.Mutex类型的互斥量，保护map访问
)

// 函数入参为字符串类型的key，从map中取值后返回。本函数用于在并发环境中使用，保证并发安全
func readValue(key string) int {
	valueByKeyGuard.Lock() // 使用互斥量，对共享资源加锁

	v := valueByKey[key] // 根据键取值

	valueByKeyGuard.Unlock() // 使用互斥量，对共享资源解锁

	return v
}

// 使用defer延迟，在函数退出时并发解锁
func readvalue(key string) int {
	valueByKeyGuard.Lock()

	defer valueByKeyGuard.Unlock() // 最后执行解锁释放资源的操作。当函数返回之后，才解锁释放资源

	return valueByKey[key] // 根据键取值
}

func TestReadValue(t *testing.T) {
	readValue("abc")
}

/*
在一般的业务逻辑中，常常涉及成对操作，比如打开与关闭文件、接收与回复请求、加锁与解锁。
上述操作均需要在函数退出时，正确地释放资源。
defer语句正好是在函数退出时执行，可以很方便地释放资源。
*/
