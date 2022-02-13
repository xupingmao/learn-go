package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var a string = "中文"

	fmt.Printf("a = %v\n", a)
	fmt.Printf("len(a) = %v\n", len(a)) // 输出: len(a) = 6
	fmt.Printf("a[0] = %v\n", a[0])     // 输出: a[0] = 228

	fmt.Printf("utf8.RuneCountInString(a) = %v\n", utf8.RuneCountInString(a))
	// 输出: utf8.RuneCountInString(a) = 2

	fmt.Println(strings.Repeat("-", 30))
	var b string = "Hello,世界"
	for i:= 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(b[i:])
		// r: 是Unicode字符
		// size: 是Unicode字符utf8编码占用的字节数
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}

	fmt.Println(strings.Repeat("-", 30))
	for i,r := range b {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}