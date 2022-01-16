package main

import "fmt"
import "tempconv"

// 需要把tempconv的实现移动到$GOPATH路径里面，无法相对路径引用
// 004_package.go:4:8: cannot find package "tempconv" in any of:
//	/usr/local/Cellar/go/1.10.1/libexec/src/tempconv (from $GOROOT)
//	/Users/xupingmao/go/src/tempconv (from $GOPATH)

// 包的结构如下（需要新建一个目录作为包名，然后里面放go语言的实现，
//             源代码的文件名无要求，需要代码声明package）
// ~/go/src/tempconv/*.go

func main() {
	// struct {
	// 	name string;
	// 	age int;
	// } mike := {"Mike", 20}

	// fmt.Printf("Mike = %v\n", mike)

	var c1 tempconv.Celsius
	var c2 tempconv.Celsius = 25.5
	var c3 = tempconv.Celsius(26.5)
	c4 := tempconv.Celsius(27.5)

	var f1 tempconv.Fahrenheit = 100.5

	fmt.Printf("c1 = %v, type: %T\n", c1, c1)
	fmt.Printf("c1.ToString() = %v\n", c1.ToString())

	fmt.Println()
	fmt.Printf("c2 = %v, type: %T\n", c2, c2)
	fmt.Printf("c2.ToString() = %v\n", c2.ToString())

	fmt.Println()
	fmt.Printf("c3.ToString() = %v\n", c3.ToString())
	fmt.Printf("c3.CToF() = %v\n", c3.CToF())

	fmt.Println()
	fmt.Printf("c4.ToString() = %v\n", c4.ToString())
	fmt.Printf("c4.CToF() = %v\n", c4.CToF())

	fmt.Println()
	fmt.Printf("f1.ToString() = %v\n", f1.ToString())
	fmt.Printf("f1.FToC() = %v\n", f1.FToC())
}

