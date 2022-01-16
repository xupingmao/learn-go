
package main

import "fmt"

func switch1(arg string) {
	// 第一种模式：取值匹配
	// switch的条件类型必须和case后的数据类型匹配
	switch arg {
	case "name":
		fmt.Printf("arg.name\n")
	case "age":
		fmt.Printf("arg.age\n")
	default:
		fmt.Printf("arg.other\n")
	}
}

func switchTagless(arg int) {
	// 第二种模式：逐个判断
	fmt.Printf("Score: %v, ", arg)
	switch {
	case arg < 60:
		fmt.Printf("Bad\n")
	case arg >= 60 && arg < 80:
		fmt.Printf("Normal\n")
	case arg >= 80:
		fmt.Printf("Good\n")
	}
}

func switch3(arg int) {
	fmt.Printf("Score: %v, ", arg)
	switch true {
	case arg < 60:
		fmt.Printf("Bad\n")
	case arg >= 60 && arg < 80:
		fmt.Printf("Normal\n")
	case arg >= 80:
		fmt.Printf("Good\n")
	}
}


func main() {
	fmt.Printf("\nSwitch Normal:\n")
	switch1("name")
	switch1("skin")

	fmt.Printf("\nSwitch tagless:\n")
	switchTagless(10)
	switchTagless(60)
	switchTagless(70)
	switchTagless(95)

	fmt.Printf("\nSwitch Case 3\n")
	switch3(20)
}