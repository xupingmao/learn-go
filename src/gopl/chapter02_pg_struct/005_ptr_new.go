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
	mikeP.age = 22
	fmt.Printf("mike = %v, type:%T\n", mike, mike)
	fmt.Printf("mikeP = %v, type:%T\n", mikeP, mikeP)
	fmt.Printf("mark = %v, type:%T\n", mark, mark)
}