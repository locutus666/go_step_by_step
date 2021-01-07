package goto_test

import (
	"fmt"
	"os"
	"testing"
)

func firstCheckError(msg string) error {
	if _, err := fmt.Println(msg); err != nil {
		return err
	}

	return nil
}

func secondCheckError(msg string) error {
	if _, err := fmt.Print(msg); err != nil {
		return err
	}

	return nil
}

func exitProcess() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}

// 使用传统的编码方式，连续退出多层循环
func TestExitLoop(t *testing.T) {
	var breakAgain bool

	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时，退出循环
			if y == 2 {
				// 设置退出标记
				breakAgain = true // 默认情况下，循环只能一层一层退出，为此我们就需要设置一个状态变量breakAgain。当需要循环退出时，设置这个变量为true。

				// 退出本次内层循环
				break
			}
		}

		// 根据标记，还需要退出一次外层循环
		if breakAgain {
			break
		}
	}

	fmt.Println("done")
}

// 使用goto退出多层循环
// 使用goto语句后，无须额外的变量就可以快速退出所有循环
func TestGotoExitLoop(t *testing.T) {

	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时，退出循环
			if y == 2 {
				// 跳转到标签
				goto breakHere // 使用goto语句跳转到指定的标签
			}
		}

	}

	// 手动返回，避免代码进入标签中执行
	return

	// 声明标签，标签只能被goto使用
breakHere:
	fmt.Println("done")
}

// 使用传统的编码方式，集中处理错误
func TestCheckError(t *testing.T) {
	err := firstCheckError("I can't fly this bike forever!")
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}

	err = secondCheckError("I can fly this bike forever!")

	// 重复的错误处理代码。后期，如果陆续在这些代码中添加更多判断，就需要在每一块雷同代码中依次修改，这样极易造成疏忽和错误。
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}

	fmt.Println("done")
}

// 使用goto集中处理错误
func TestGotoFor(t *testing.T) {
	err := firstCheckError("I can't fly this bike forever!")
	if err != nil {
		goto onExit // 发生错误时，跳转到错误标签onExito
	}

	err = secondCheckError("I can fly this bike forever!")
	if err != nil {
		goto onExit // 发生错误时，跳转错误标签onExito
	}

	fmt.Println("done")

	return

	// 定义标签，汇总所有流程进行错误打印井退出进程
onExit:
	fmt.Println(err)
	exitProcess()
}

/*
使用goto语句，可以简化一些代码的实现过程。

通过使用标签，goto语句可以进行代码间的无条件跳转。

在快速跳出循环、避免重复退出等问题上，使用goto语句有一定的优势。
*/
