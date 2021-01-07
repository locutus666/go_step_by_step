package struct_test

import (
	"testing"
)

type BasicColor struct {
	R, G, B float32
}

type Color struct {
	BasicColor
	Alpha float32
}

func TestNeiqianStruct(t *testing.T) {
	var c Color

	c.R = 1
	c.G = 1
	c.B = 0

	c.Alpha = 1

	t.Logf("%+v", c)
}
