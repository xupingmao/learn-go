package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map是Go语言的关键字
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// input.Err() 可以读取错误信息
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 运行:
// go run 003_dup1.go < 003_dup1.txt
