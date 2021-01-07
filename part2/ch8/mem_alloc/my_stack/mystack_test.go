package stack_test

import "testing"

func calc(a, b int) int {
	var x, c int

	c = a + b
	x = c * 10

	return x
}

func TestMystack(t *testing.T) {
	y := calc(1, 2)
	t.Log(y)
}
