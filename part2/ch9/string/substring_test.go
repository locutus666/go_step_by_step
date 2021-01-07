package string_test

import (
	"fmt"
	"strings"
	"testing"
)

// 返回从第二个死神开始的子串
func TestSubstring(t *testing.T) {
	tracer := "死神来了，死神byebye"
	comma := strings.Index(tracer, "，") // 使用strings.Index，在字符串"死神来了，死神byebye"中搜索子串"，"，返回"，"位置的ascii字符串下标

	pos := strings.Index(tracer[comma:], "死神") // 从字符串"死神来了，死神byebye"中的"，"位置开始，一直到字符串结束，构造一个子串"，死神byebye"，在这个子串中搜索"死神"，返回下标（相对位置的偏移）

	fmt.Println(comma, pos, tracer[comma+pos:]) // 输出"，"所在字符串下标，"死神"所在子串的下标，利用切片偏移（绝对位置+相对位置）输出子串
}

// 12 3 死神byebye
