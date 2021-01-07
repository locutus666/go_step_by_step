package go_test

import (
	"fmt"
	"testing"
)

func Benchmark_Add1(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("%d", i)
	}
}
