package interface_test

import (
	"fmt"
	"testing"
)

// 接口定义
type Flyer interface { // 定义飞行动物接口
	Fly()
}

type Walker interface { // 定义行走动物接口
	Walk()
}

// 接口实现
type bird struct{} // 定义飞行类型

func (b *bird) Fly() { // 实现飞行动物飞行接口
	fmt.Println("bird: fly")
}

func (b *bird) Walk() { // 实现飞行动物行走接口
	fmt.Println("bird: walk")
}

// 接口实现
type pig struct{} // 定义行走类型

func (p *pig) Walk() { // 实现行走动物行走接口
	fmt.Println("pig: walk")
}

func TestBirdPig(t *testing.T) {
	// 创建动物名到结构体实例的字典
	animals := map[string]interface{}{
		"bird": new(bird), // 创建出的结构体实例
		"pig":  new(pig),
	}

	// 遍历字典
	for name, obj := range animals { // obj为字典的值，是interface{}类型
		f, isFlyer := obj.(Flyer)   // 使用类型断言获得变量f，转换后的类型是Flyer；isFlyer是接口类型转换是否成功的结果
		w, isWalker := obj.(Walker) // 使用类型断言获得变量w，转换后的类型是Walker；isWalker是接口类型转换是否成功的结果

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly() // 调用接口方法
		}

		if isWalker {
			w.Walk() // 调用接口方法
		}
	}
}

/*
name: bird isFlyer: true isWalker: true
bird: fly
bird: walk
name: pig isFlyer: false isWalker: true
pig: walk
*/
