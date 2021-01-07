package if_test

import (
	"fmt"
	"testing"
)

func TestIfElseifElse(t *testing.T) {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 { // else if与if的条件判断同级
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
