package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer func() {
		// kill -9 会执行
		fmt.Println("defer executed")
		// kill -9 下recover() 返回 nil
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
		}
	}()

	var sc = bufio.NewScanner(os.Stdin)
	fmt.Printf("请输入:")
	// 等待输入
	sc.Scan()

	// >>> kill -9 输出结果
	// 请输入:defer executed
	// Killed: 9

	// >>> Ctrl+C终止
	// 请输入:^Csignal: interrupt
}
