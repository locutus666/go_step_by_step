package string_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestSubstringSearch(t *testing.T) {
	tracer := "死神来了，死神byebye"

	// 使用strings.Index，在字符串"死神来了，死神byebye"中搜索子串"，"，返回"，"在tracer字符串中位置的ascii下标
	comma := strings.Index(tracer, "，")
	fmt.Println(comma)
}

// 12
