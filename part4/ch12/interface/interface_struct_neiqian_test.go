package interface_test

// 接口定义方
type Service interface {
	Start()
	Log(string)
}

// 接口实现方
type Logger struct{}

func (g *Logger) Log(l string) { // 实现Service接口的Log方法
}

type GameService struct { // 结构体内嵌
	Logger
}
