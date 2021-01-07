package struct_test

import (
	"fmt"
	"testing"
)

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) IntSlice {
	p[i], p[j] = p[j], p[i]
	return p
}

func TestTypeBindFunc(t *testing.T) {
	var a IntSlice

	a = []int{1, 2, 3}

	fmt.Println(a.Len())
	fmt.Println(a.Less(1, 2))
	fmt.Println(a.Swap(1, 2))
}
