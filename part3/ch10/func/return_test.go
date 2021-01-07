package func_test

import (
	"testing"
)

func div(a, b int) (c int, d int) {
	return a / b, a % b
}

func TestReturn(t *testing.T) {
	q, _ := div(17, 3)
	t.Log(q)
}
