package main

import (
	"time"

	"github.com/pkg/profile" // 在设置好GOPATH条件下，第三方包需要使用go get下载到本地才能用，引用路径从$GOPATH/src/之后开始
)

func joinSlice() []string {
	var arr []string

	for i := 0; i < 100000; i++ {
		arr = append(arr, "arr") // 向arr切片中，不停的添加元素。o(n)，性能较低
	}
	return arr
}

func main() {
	// 使用profile.Start调用github.com/pkg/profile开始性能分析，返回一个stopper.stop接口，方便在程序结束时，结束性能分析
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath(".")) // 这里使用profile.CPUProfile（CPU耗用）作为分析项，使用profile.ProfilePath(".")指定输出程序性能分析报告到当前路径

	// 在main()函数结束时，结束性能分析
	defer stopper.Stop()

	// 性能分析核心逻辑
	joinSlice()

	// 程序性能分析最短时间为1秒，在main()函数结束前等待1秒，如果你的程序默认运行1秒以上，可以去掉这句
	time.Sleep(time.Second)
}
