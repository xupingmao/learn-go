package main

import "fmt"

type Person struct {
	name string
	age int
}

func changeValue(p *int) {
	*p += 1
}

func main() {
	x := 1
	p := &x

	mike := Person{"Mike", 21}
	bill := Person{name:"Bill", age:23}

	// new函数只是一个语法糖，和初始化变量之后取址是等价的
	// 等价于
	// var instance Person;
	// mark := &instance;
	mark := new(Person)
	*mark = Person{"Mark", 35}

	fmt.Printf("x = %v, type:%T\n", x, x)
	fmt.Printf("p = %v, type:%T\n", p, p)

	changeValue(p)
	fmt.Printf("\nAfter value changed...\n")
	fmt.Printf("x = %v, type:%T\n", x, x)
	fmt.Printf("p = %v, type:%T\n", p, p)

	fmt.Printf("mike = %v, type:%T\n", mike, mike)
	fmt.Printf("bill = %v, type:%T\n", bill, bill)

	mikeP := &mike
	mikeP.age = 22 // 这个和C语言的语法不太一样，直接用点来访问就可以了，不需要使用->
	fmt.Printf("mike = %v, type:%T\n", mike, mike)
	fmt.Printf("mikeP = %v, type:%T\n", mikeP, mikeP)
	fmt.Printf("mark = %v, type:%T\n", mark, mark)
}