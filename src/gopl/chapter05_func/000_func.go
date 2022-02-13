package main


import (
	"fmt"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

type Number int;
type BigNumber int64;

func (p Number) Add(other Number) Number {
	return p + other;
}

func (p BigNumber) Add(other Number) Number {
	return Number(p) + other;
}

// 下面函数报错: cannot define new methods on non-local type string
// func (str string) Add(other string) string {
// 	return str + "," + other;
// }

// 必须定义一个新的类型
type MyString string;
func (str MyString) Add(other MyString) MyString {
	return str + "," + other;
}

func funcAsParam(f func(Number)Number, param Number) {
	var result Number = f(param);
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("func: %T\n", f);
	fmt.Printf("funcAsParam: result:%v\n", result);
}

func main() {
	fmt.Printf("%T, result:%v\n", add, add(1,2))
	fmt.Printf("%T, result:%v\n", sub, sub(10,5))
	fmt.Printf("%T, result:%v\n", first, first(5,10))
	fmt.Printf("%T, result:%v\n", zero, zero(1,2))

	var num1 = Number(1);
	var num2 = Number(2);
	fmt.Printf("%T, result:%v\n", num1.Add(num2), num1.Add(num2));

	var bigNum1 = BigNumber(1);

	// 方法的签名上不带类的信息
	// Number.Add 和 BigNumber.Add 是同一个方法签名
	// 那么类型的值或者指针是如何传入的呢？
	funcAsParam(num1.Add, num2);
	funcAsParam(bigNum1.Add, num2);

	var a MyString = "Hello"
	var b MyString = "World"
	fmt.Printf("a+b=%v\n", a.Add(b));
}