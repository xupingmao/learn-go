package main

import "fmt"

func main() {
	// 下面语句报错: const declaration cannot have type without expression
	// const a int
	const b1 int = 1
	// 类型可以省略，编译器可以自动推导为默认类型
	const b2 = 1
	// 如果需要其他类型，需要声明出来
	const b3 int64 = 1

	// 下面报错: missing value in const declaration
	// const d, e, f int

	// fmt.Printf("a = %v, type:%T\n", a, a)
	fmt.Printf("b1 = %v, type:%T\n", b1, b1)
	fmt.Printf("b2 = %v, type:%T\n", b2, b2)
	fmt.Printf("b3 = %v, type:%T\n", b3, b3)
	// fmt.Printf("c = %v, type:%T\n", c, c)
	// fmt.Printf("d = %v, type:%T\n", d, d)
	// fmt.Printf("e = %v, type:%T\n", e, e)
	// fmt.Printf("f = %v, type:%T\n", f, f)
}

