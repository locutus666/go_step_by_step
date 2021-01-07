package map_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestMapDelete(t *testing.T) {
	// 创建map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	// 打印map
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 删除元素
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	// 打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

// 添加key:value到map，并查找map中key对应的value，遍历map中的键值对
func TestMapAddVisit(t *testing.T) {
	scene := make(map[string]int) // 使用make手动创建map，赋值给变量scene
	t.Log(scene)
	// var scene map[string]int // 如果不创建就是用map，会触发panic: assignment to entry in nil map [recovered]

	scene["route"] = 66 // 向map中添加映射关系，key可以是任何类型（函数类型除外）
	scene["brazil"] = 4
	scene["china"] = 960
	t.Log(scene)
	t.Log(scene["route"]) // 查找map中对应key的值

	v := scene["route2"] // 查找map中一个不存在的key，会返回map中定义value_data_type的零值
	t.Log(v)

	v, ok := scene["route"] // 在获取键值v的基础上，多取一个变量ok，判断map中是否存在key route
	t.Log(ok)               // 若key存在，则ok为true

	// 使用for range，迭代map，遍历键值对
	for k, v := range scene { // 直接使用for range遍历，同时获得key、value
		t.Log(k, v)
	}

	for _, v := range scene { // 把不需要的键改为匿名变量_的形式
		t.Log(v)
	}

	for k := range scene { // 只遍历键，无须把值改为匿名变量_的形式，忽略v即可
		t.Log(k)
	}

}

//
func TestMapSort(t *testing.T) {
	scene := make(map[string]int)
	t.Log(scene)

	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	t.Log(scene)

	// 声明一个切片，用于缓冲和排序map元素
	var sceneList []string

	// 把map键k遍历，并复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}

	// 对切片中的字符串元素进行排序
	sort.Strings(sceneList)
	t.Log(sceneList) // 输出排序后的map键

	for _, k := range sceneList {
		t.Log(scene[k])
	}
}

// sort.Strings()对字符串字符进行升序排列
