package string_test

import (
	"testing"
)

// 字符串与字符编码，Unicode与UTF-8
func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值""

	s = "hello"
	t.Log(len(s)) // len(string)返回对应字符的byte数

	// s[1] = '3' //string是不可变的byte slice，对string赋值会报错cannot assign to s[1]
	// t.Log(len(s))

	s = "\xE4\xB8\xA5" //string可以存储二进制编码数据（16进制），如汉字“严”
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xBA\xBB\xFF" // string可以存储任意的二进制编码数据（16进制），虽然显示乱码
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s)) //中文字符长度是其对应的byte数，这里“中”字对应3个byte

	c := []rune(s) // 使用rune类型的切片，取出string中的unicode。一个rune类型的数据代表其对应的Unicode编码
	t.Log(c)       // 返回含“中”对应unicode值的rune类型的切片
	t.Log(len(c))  // 切片长度是1
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0]) // 输出切片c第一个元素的十六进制表示，即“中”字的unicode编码是ox4e2d，占2byte
	t.Logf("中 UTF8 %x", s)       // 输出“中”字的UTF8物理存储字节序列，即“中”字的utf8是oxe4b8ad，占3byte。存放在[]byte的切片[0xE4, 0xB8, 0xAD]中，每个byte各占一个位置
}

// 使用for range，遍历unicode字符串
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s { // string遍历的第一个值是字符的序号（从0开始），可以用"_"舍弃
		t.Log(c)                        // 默认返回字符串中每个字符对应的unicode编码
		t.Logf("%[1]c %[1]d %[1]x ", c) // 返回字符串中每个字符的字符显示（%c），unicode编码（%d），十六进制数0x（%x）
		// t.Logf("%c %x", c, c) // 与上一行等价
	}
}
