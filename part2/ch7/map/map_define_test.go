package map_test

import (
	"fmt"
	"testing"
)

// 定义map方式1：使用关键字map，大括号{}，像json一样，定义key:value，键值对之间使用逗号分隔
func TestMapDefine(t *testing.T) {
	var m1 = map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)

	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m2)

	m3 := map[string]string{}
	if m3 == nil {
		fmt.Println("m3 is nil!")
	}
	t.Log(m3)

	m4 := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	t.Log(m4)
}

// 定义map方式2：使用内置函数make
func TestMap(t *testing.T) {
	countryCapitalMap := make(map[string]string)

	// map插入key - value对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	// 使用键输出map值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看key在map中是否存在
	capital, ok := countryCapitalMap["美国"] // 如果确定是真实的,则存在,否则不存在
	// fmt.Println(capital)
	// fmt.Println(ok)
	if ok {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("美国的首都不存在")
	}
}
