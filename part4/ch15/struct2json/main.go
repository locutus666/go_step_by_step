package main

import "fmt"

func main() {

	// 声明技能结构
	type Skill struct {
		Name  string
		Level int
	}

	// 声明角色结构
	type Actor struct {
		Name string
		Age  int

		Skills []Skill
	}

	// 填充角色的基本数据
	a := Actor{
		Name: "cow boy",
		Age:  37,

		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	if result, err := MarshalJSON(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 声明技能结构
	type Skill struct {
		Name  string
		Level int
	}

	// 声明角色结构
	type Actor struct {
		Name string
		Age  int

		Skills []Skill
	}

	// 填充角色的基本数据
	a := Actor{
		Name: "cow boy",
		Age:  37,

		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	if result, err := json.Marshal(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
*/
