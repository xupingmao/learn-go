package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func testIndexOutOfRange() {
	arr2 := [3]int{1,2,3}

	sc:=bufio.NewScanner(os.Stdin)
	fmt.Print("index:")
	// 读取输入值
	sc.Scan()
	index_str := sc.Text()
	index, _ := strconv.ParseInt(index_str, 10, 64)

	arr2[index] = 10
	fmt.Printf("arr = %o\n", arr2);

	// 输入10之后程序崩溃
	// panic: runtime error: index out of range
}


func main() {
	arr := [2]int{1,2}
	fmt.Printf("arr = %o\n", arr);

	testIndexOutOfRange()
}