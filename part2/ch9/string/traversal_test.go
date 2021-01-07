package string_test

import (
	"fmt"
	"testing"
)

func TestStringTraversal(t *testing.T) {
	theme := "狙击 start"

	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}
}

func TestForRangeString(t *testing.T) {
	theme := "狙击 start"

	for _, s := range theme {
		fmt.Printf("ascii: %c %d\n", s, s)
	}
}

/*
ascii: ç 231
ascii:  139
ascii:  153
ascii: å 229
ascii:  135
ascii: » 187
ascii:   32
ascii: s 115
ascii: t 116
ascii: a 97
ascii: r 114
ascii: t 116

ascii: 狙 29401
ascii: 击 20987
ascii:   32
ascii: s 115
ascii: t 116
ascii: a 97
ascii: r 114
ascii: t 116
*/
