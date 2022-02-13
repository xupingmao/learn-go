package main

import (
	"fmt"
	"strings"
)

type Address struct {
	province string
	detail string
}

type Person struct {
	name string
	age int
	addressList []Address
}

func (p Person) Print(tag string) {
	fmt.Printf("Print: [%v] name:%v, age:%v\n", tag, p.name, p.age)
}

func (p Person) UpdateName(name string) {
	// 这里的p是复制的，修改这里的值不会影响传入的参数
	p.name = name;
	fmt.Printf("UpdateName: newName:%v\n", name)
}

func (p *Person) Print2(tag string) {
	fmt.Printf("Print2: [%v] name:%v, age:%v\n", tag, p.name, p.age)
}

func (p *Person) UpdateName2(name string) {
	p.name = name;
	fmt.Printf("UpdateName2: newName:%v\n", name)
}


func main() {
	var p Person = Person{name: "Mark", age: 20, addressList: []Address {{
		province: "浙江", detail:"浙江省西湖区"}}}

	var pCopy = p;

	fmt.Printf("p = %v\n", p)


	var ptr *Person = &p;

	p.Print("p");
	// 自动取指针
	p.Print2("p")
	p.UpdateName("Jack") // 调用修改后p.name还是Mark
	p.Print("p")

	fmt.Println(strings.Repeat("-", 50))

	// 自动加解指针
	ptr.Print("ptr")
	ptr.Print2("ptr")
	ptr.UpdateName2("Jack") // 调用修改后p.name会变成Jack
	ptr.Print("ptr")

	fmt.Println(strings.Repeat("-", 50))
	pCopy.UpdateName2("Jack") // 这里会自动转指针，也可以修改pCopy的值
	pCopy.Print("pCopy") // 修改后pCopy.name变成Jack
}