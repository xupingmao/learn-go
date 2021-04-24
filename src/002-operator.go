
package main

import "fmt"


func TestMathOperator() {
	var a = 10;
	var b = 20;

	fmt.Printf("a + b = %d\n", a + b)
	fmt.Printf("a - b = %d\n", a - b)
	fmt.Printf("a * b = %d\n", a * b)
	fmt.Printf("a / b = %d\n", a / b)
	fmt.Printf("a %% b = %d\n", a % b)
}

func TestLogicOperator() {
	var a = 10;
	var b = 20;

	if (a < b) {
		fmt.Printf("a is litter than b\n")
	}
}

func main() {
	TestMathOperator()
	TestLogicOperator()
}