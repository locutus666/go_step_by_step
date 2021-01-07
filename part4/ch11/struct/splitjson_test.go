package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int // 电池容量
}

// 生成json数据
func genJsonData() []byte {

	raw := &struct { // 实例化
		Screen // 结构体内嵌
		Battery
		HasTouchID bool
	}{
		Screen: Screen{ // 初始化成员变量
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		Battery: Battery{
			2910,
		},

		HasTouchID: true,
	}

	// 把数据序列化为json
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

// 分离json数据
func TestSplitJson(t *testing.T) {
	// 生成一段json数据
	jsonData := genJsonData()
	fmt.Println(string(jsonData))

	// 只需要屏幕和指纹识别信息
	screenAndTouch := struct { // 定义匿名结构体
		Screen     // 结构体内嵌
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息
	batteryAndTouch := struct { // 定义匿名结构体
		Battery    // 结构体内嵌
		HasTouchID bool
	}{}

	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)
}
