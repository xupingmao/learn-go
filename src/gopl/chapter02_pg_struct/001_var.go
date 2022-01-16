package main

import "fmt"

func main() {
	// 变量的三种声明方式
	// var name type = expression
	var a int
	var b1 int = 1
	// 类型可以省略，编译器可以自动推导为默认类型
	var b2 = 1
	// 如果需要其他类型，需要声明出来
	var b3 int64 = 1
	c := 2

	// 多个变量声明
	var d, e, f int

	fmt.Printf("a = %v, type:%T\n", a, a)
	fmt.Printf("b1 = %v, type:%T\n", b1, b1)
	fmt.Printf("b2 = %v, type:%T\n", b2, b2)
	fmt.Printf("b3 = %v, type:%T\n", b3, b3)
	fmt.Printf("c = %v, type:%T\n", c, c)
	fmt.Printf("d = %v, type:%T\n", d, d)
	fmt.Printf("e = %v, type:%T\n", e, e)
	fmt.Printf("f = %v, type:%T\n", f, f)
}