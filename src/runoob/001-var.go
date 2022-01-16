
package main

import "fmt"


// Go 语言中变量可以在三个地方声明：
// - 函数内定义的变量称为局部变量
// - 函数外定义的变量称为全局变量
// - 函数定义中的变量称为形式参数


var (
	VAR1 int
	VAR2 string
)

func PrintFuncHeader(funcName string) {
	fmt.Println("===========================")
	fmt.Printf("(%s)\n", funcName)
}


func AutoTypeNumber() {
	PrintFuncHeader("AutoTypeNumber")
	
	var age = 10;
	fmt.Printf("my age is %d\n", age)
}

func AutoTypeString() {
	PrintFuncHeader("AutoTypeString")

	var name = "Mike";
	fmt.Printf("my name is %s\n", name)

	var country string = "China";
	fmt.Printf("I come from %s\n", country)

	fruit := "apple"
	fmt.Printf("%s is my favorite fruit\n", fruit)
}

func DefType() {
	PrintFuncHeader("DefType")

	var age int;
	age = 10;
	fmt.Printf("my age is %d\n", age)
}

func DefTypeGlobal() {
	PrintFuncHeader("DefTypeGlobal")

	VAR1 := 10
	VAR2 := "test"

	// 声明只能进行一次，这里会报错
	// ./001-var.go:39:7: no new variables on left side of :=
	// VAR1 := 20
	fmt.Printf("VAR1  = %d\n", VAR1)
	fmt.Printf("VAR2  = %s\n", VAR2)
	fmt.Printf("&VAR1 = %p\n", &VAR1)
}

func DefTypeGlobal2(GLOBAL_VAR1 int) {
	PrintFuncHeader("DefTypeGlobal2")
	
	// 这里不报错了，其实这里的VAR1和全局变量VAR1是两个变量
	VAR1 := 20

	fmt.Printf("VAR1         = %d\n", VAR1)
	fmt.Printf("&VAR1        = %p\n", &VAR1)
	fmt.Printf("&GLOBAL_VAR1 = %p\n", &GLOBAL_VAR1)
}

func main() {
	AutoTypeNumber();
	AutoTypeString();

	DefType();
	DefTypeGlobal();
	DefTypeGlobal2(VAR1);
}
