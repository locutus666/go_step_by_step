package ptr_test

import (
	"fmt"
	"testing"
)

func TestPtradrestype(t *testing.T) {
	var cat int = 1

	var str string = "orange"

	fmt.Printf("%p %p", &cat, &str)
}

// 0xc0000b8030
// 0xc0000ae030
