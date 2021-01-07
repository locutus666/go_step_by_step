package string_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStringJo(t *testing.T) {
	hammer := "吃我一锤，"
	sickle := "去死吧！"

	he := hammer + sickle
	fmt.Println(he)
}

func TestStringJoin(t *testing.T) {
	hammer := "吃我一锤，"
	sickle := "去死吧！"

	// 声明字节缓冲。bytes.Buffer可以缓冲并可以往里面写入各种字节数组，字符串也是一种字节数组，使用WriteString()方法进行写入。
	var stringBuilder bytes.Buffer

	// 把需要连接的字符串，通过调用WriteString()方法，写入stringBuilder中
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	// 再通过stringBuilder.String()方法将缓冲转换为字符串。
	fmt.Println(stringBuilder.String())
}
