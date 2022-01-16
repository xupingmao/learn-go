package main

import "fmt"

// type语法
// type name underlying-type

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToString() string {
	return fmt.Sprintf("%gC", c)
}

func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit((c*9/5+32))
}

func (f Fahrenheit) ToString() string {
	return fmt.Sprintf("%gF", f)
}

func (f Fahrenheit) FToC() Celsius {
	return Celsius((f-32)*5/9)
}

func main() {
	// struct {
	// 	name string;
	// 	age int;
	// } mike := {"Mike", 20}

	// fmt.Printf("Mike = %v\n", mike)

	var c1 Celsius
	var c2 Celsius = 25.5
	var c3 = Celsius(26.5)
	c4 := Celsius(27.5)

	var f1 Fahrenheit = 100.5

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

