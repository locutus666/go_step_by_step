package slice_test

import (
	"testing"
)

// 切片扩容为原容量的2倍：切片一次新增元素数不超过原容量的1倍，且不超过1024个，且增加元素后新切片总长度小于1024个
func TestSliceCap1(t *testing.T) {
	s1 := make([]int, 0) // 初始化一个空切片[]
	t.Log(s1)
	t.Logf("The capacity of s1: %d\n", cap(s1))
	t.Log("")

	for i := 1; i <= 17; i++ {
		s1 = append(s1, i)
		t.Log(s1)
		t.Logf("s1(%d): len: %d, cap: %d\n", i, len(s1), cap(s1))
		t.Log("")
	}
}

// 切片扩容为原容量的2倍稍大一些：切片一次新增元素数超过原容量的1倍，但不超过1024个，且增加元素后新切片总长度小于1024个
func TestSliceCap2(t *testing.T) {
	s2 := make([]int, 10) // 初始化一个切片，用类型零值填充[0 0 0 0 0 0 0 0 0 0]
	t.Log(s2)
	t.Logf("The capacity of s2: %d\n", cap(s2)) // 原切片cap = 10
	t.Log("")

	r1 := append(s2, make([]int, 5)...)
	t.Logf("r1: len: %d, cap: %d\n", len(r1), cap(r1))

	r2 := append(s2, make([]int, 11)...)
	t.Logf("r2: len: %d, cap: %d\n", len(r2), cap(r2))

	r3 := append(s2, make([]int, 21)...)
	t.Logf("r3: len: %d, cap: %d\n", len(r3), cap(r3))

	t.Logf("注意：像r2，r3一次增加元素的个数（11，21）超过原切片容量（10）的1倍，新切片容量比原切片的2倍稍大一些\n")
	t.Log("")
}

// 切片扩容为原容量的1.25倍：原切片长度大于等于1024时，一次增加元素个数不超过原切片长度的0.25倍，在一次扩容后，新切片容量为原切片容量的1.25倍
func TestSliceCap3(t *testing.T) {
	// 1024 * 0.25 = 256
	s3 := make([]int, 1024)
	t.Log(s3)
	t.Logf("The capacity of s3: %d\n", cap(s3))
	t.Log("")

	r4 := append(s3, make([]int, 200)...)
	t.Logf("r4: len: %d, cap: %d\n", len(r4), cap(r4))

	r5 := append(s3, make([]int, 256)...)
	t.Logf("r5: len: %d, cap: %d\n", len(r5), cap(r5))

	t.Logf("注意：像r4，r5一次增加元素数不超过原切片容量的0.25倍，扩容后，新切片容量是原切片的1.25倍\n")
	t.Log("")
}

// 切片扩容为原容量的1.25倍稍大一些：原切片长度大于等于1024时，一次增加元素个数超过原切片长度的0.25倍，在一次扩容后，新切片容量比原切片容量的1.25倍稍大一些
func TestSliceCap4(t *testing.T) {
	// 1024 * 0.25 = 256
	s4 := make([]int, 1024)
	t.Log(s4)
	t.Logf("The capacity of s4: %d\n", cap(s4))
	t.Log("")

	// 1280 * 0.25 = 320
	r6 := append(s4, make([]int, 266)...)
	t.Logf("r6: len: %d, cap: %d\n", len(r6), cap(r6))

	t.Logf("注意：像r6一次增加元素数超过原切片容量的0.25倍，扩容后，新切片容量比原切片的1.25倍稍大一些\n")
	t.Log("")
}

/*
比如r6, 一次增加元素之后，新切片长度大于原切片容量的1.25倍（1024+266=1290 > 1024*1.25=1280），需要扩容。若按原切片容量的1.25倍扩容，则1280*1.25=1600，但是实际扩容量（1696）比原切片的1.25倍累乘（1024*1.25*1.25 = 1600）稍大一些。
*/
