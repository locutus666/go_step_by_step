package func_test

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer fmt.Println("宕机后要做的事1")
	defer fmt.Println("宕机后要做的事2")
	panic("宕机")
}
