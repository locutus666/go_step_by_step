package ptr_test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	str := new(string)
	*str = "ninja"

	fmt.Println(*str)
}
