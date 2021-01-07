package string_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// 根据文件名，段名，键名获取ini的值
func getValue(filename string, expectSection string, expectKey string) string {

	//// 使用文件打开函数os.Open()，以二进制的方式打开文件。go总是以二进制的方式打开文件。
	// 如果打开文件成功，返回文件句柄file和可能出现的错误err，err为nil，函数继续往下执行
	file, err := os.Open(filename)

	// 如果打开文件错误，err不为nil，函数返回一个空字符串，表示无法从.ini文件中读取内容，最后执行defer语句
	if err != nil {
		return ""
	}

	// 在函数getValue结束时，即return返回值后，关闭文件并释放系统资源
	defer file.Close()

	//// go可以使用不同的读取方式，操作打开的二进制文件。读取.ini文本时，需要注意各种异常情况。
	// 使用函数bufio.NewReader，构造一个读取器reader，读取文件句柄file，获取行文本
	reader := bufio.NewReader(file)

	// 构造一个读取循环，不断读取文件的每一行
	for {
		// 使用reader.ReadString()从打开的二进制文件中读取行文本，返回读取到的行字符串linestr（包括 "\n"）和可能读取的错误err，直到碰到“＼n”，也就是行结束。
		linestr, err := reader.ReadString('\n')

		// 如果读取错误，break跳出循环
		if err != nil {
			break
		}

		// 切掉行左右两边的空白字符。
		// 每一行的文本可能会在标识符两边混杂有空格、回车符、换行符等不可见的空白字符，使用strings.TrimSpace()可以去掉这些空白字符。
		linestr = strings.TrimSpace(linestr)

		// 忽略空行。
		// 可能存在空行的情况，忽略空行，continue继续读取下一行。
		if linestr == "" {
			continue
		}

		// 忽略注释。
		// 当行首的字符为“；”分号时，表示这一行是注释行，使用continue关键字，忽略这一行的读取
		if linestr[0] == ';' {
			continue
		}

		//// 抛开各种异常情况拿到了每行的行文本linestr后，就可以方便地读取.ini文件的段和键值了。
		// 声明读取的段名变量
		var sectionName string

		// 如果行首和行尾分别是方括号，说明是段名的起止符
		// linestr[O]表示行首字符，len(linestr)-1表示取出字符串的最后－个字符索引，再取出行尾字符。根据两个字符串是否同时匹配方括号，判断当前行是否为段名。
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {

			// 把段名取出。使用索引，把linestr两边的"["和"]"去掉，取出中间的段名，保存在sectionName中，留着后面代码用。
			sectionName = linestr[1 : len(linestr)-1]
			fmt.Println(sectionName)

		} else if sectionName == expectSection { // 如果行字符串linestr不是段名，则把我们希望读取的段名expectSection，赋给段名变量

			// 切开等号分割的键值对。
			// 通过strings.Split()函数，对行内容linestr进行切割，分隔符为“=”。切割后，strings.Split()函数会返回含两个元素的字符串切片，类型为[]string，即pair，pair[0]是键，pair[1]是值，
			pair := strings.Split(linestr, "=")

			// 只考虑切割出2个元素的情况，保证切开的键值只有1个等号分割
			if len(pair) == 2 {

				// 去掉键的多余空白字符
				key := strings.TrimSpace(pair[0])

				// 切割键值对后，还需要判断键是否为期望的键
				if key == expectKey {

					// 匹配到期望的键时，把pair中保存键对应的值，经过去掉空白字符处理后，作为函数返回值返回
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}

	// 如果读取错误，break跳出循环，返回空字符串；
	return ""
}

func TestStringOpen(t *testing.T) {
	t.Log(getValue("example.ini", "remote \"origin\"", "fetch"))

	t.Log(getValue("example.ini", "core", "hideDotFiles"))
}

// getValue这个函数的作用是读取.ini配置文件，提取段名，提取期望段指定的键值
