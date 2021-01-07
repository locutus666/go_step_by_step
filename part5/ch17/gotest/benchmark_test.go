package go_test

// package benchmark_test，在同一个包文件夹下，go源文件的package声明必须是同一个包名，否则会报错

import "testing"

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
