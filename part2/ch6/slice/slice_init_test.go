package slice_test

import "testing"

// 初始化声明slice
func TestSliceInit(t *testing.T) {
	var s0 []int            // 声明空切片
	t.Log(s0[:])            // 输出切片元素
	t.Log(cap(s0), len(s0)) // 输出切片长度和容量，0 0
	s0 = append(s0, 1)      // 向切片s0添加元素1
	t.Log(cap(s0), len(s0))

	s1 := []int{} // 声明空切片
	t.Log(s1[:])
	t.Log(cap(s1), len(s1))

	s2 := []int{1, 2, 3, 4} // 初始化声明切片
	t.Log(cap(s2), len(s2)) // 输出切片长度和容量，4 4

	s3 := make([]int, 3, 4) // 声明len=3，cap=4的切片；其中len个元素会被初始化为默认零值，未初始化元素不可以访问
	t.Log(s3)
	t.Log(s3[:]) // 输出切片元素，切片的输出：s3[:]与s3一致
	// t.Log(s3[0], s3[1], s3[2], s3[3]) // len个元素会被初始化为默认零值，未初始化元素不可以访问，会引发panic: runtime error: index out of range
	t.Log(s3[0], s3[1], s3[2]) // 输出切片对应索引的元素
	s3 = append(s3, 1)
	t.Log(s3)
	t.Log(len(s3), cap(s3))
}

// 切片容量扩展
func TestSliceGrowing(t *testing.T) {
	s4 := []int{} // 切片长度和容量，0 0
	for i := 0; i < 10; i++ {
		s4 = append(s4, i)
		t.Log(len(s4), cap(s4))
	}
}

/*
使用append，在向切片插入元素后，为何还需赋值给原切片？
答：1. 切片是一个结构体，当向其中append一个元素时，切片指向的连续存储空间地址发生了变化，需要重新赋值给原切片；
   2. 当原存储空间扩容后，会创建一个新的连续存储空间，go会把原有的数值拷贝到新的连续存储空间

切片使用的优劣：切片的容量可以自增长，但是自增长的代价是需要复制原存储空间的数据到新存储空间。
*/

// 切片共享存储结构。多个slice指向同一段存储空间
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(year, len(year), cap(year)) // 12 12
	t.Log(year[0])

	Q2 := year[3:6]             // 切片截取。slice是数据结构，可以赋值给变量
	t.Log(Q2, len(Q2), cap(Q2)) // 3 9。截取的slice从索引3开始，一直指向原切片后面的连续存储空间，所以从索引3到索引11，容量为9；截取的slice长度是3

	summer := year[5:8]                     // 切片截取
	t.Log(summer, len(summer), cap(summer)) // 3 7

	summer[0] = "Unknow" // summer的第一个元素被替换
	t.Log(Q2)            // 与summer共享同一段存储空间的切片Q2、year，相应的值也随着改变
	t.Log(year)
}

// 切片只能和nil比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { //切片只能和nil比较
	// 	t.Log("equal")
	// }
	if a != nil {
		t.Log(a, b)
	}
}

/*
slice是可变长的数组，本质上是struct指针
*/
