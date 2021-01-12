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

{"Name":"cow boy","Age":37,"Skills":[{"Name":"Roll and roll","Level":1},{"Name":"Flash your dog eye","Level":2},{"Name":"Time to have Lunch","Level":3}]}

{
    "Name":"cow boy",
    "Age":37,
    "Skills":[
        {
            "Name":"Roll and roll",
            "Level":1
        },
        {
            "Name":"Flash your dog eye",
            "Level":2
        },
        {
            "Name":"Time to have Lunch",
            "Level":3
        }
    ]
}
*/

/*
以上的序列化方式，在程序功能和结构上还有一些不足：

没有处理各种异常情况，切片或结构体为空时应该提前判断，否则会触发岩机。

可以支持结构体标签（Struct Tag），方便自定义JSON的键名及忽略某些字段的序列化过程，避免这些字段被序列化到JSON中。

支持缩进且可以自定义缩进字符，将JSON序列化后的内容格式化，方便查看。

默认应该序列化为[]byte字节数组，外部自己转换为字符串。在大部分的使用中，JSON一般以字节数组方式解析、存储、传输，很少以字符串方式解析，因此避免字节数组和字符串的转换，可以提高一些性能。
*/
