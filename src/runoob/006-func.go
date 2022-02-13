package main

import (
	"fmt"
)


func Add(a,b int) int {
	return a + b
}

func Add2(a,b int) (ret int) {
	ret = a + b
	// return是必须的
	return
}

func Add3(a,b int) (ret int) {
	ret = a + b
	// 最终返回的是10
	return 10
}

func TupleAdd(a,b,inc int) (int, int) {
	return a + inc, b + inc
}

func main() {
	fmt.Printf("Add(1,2)=%v\n", Add(1,2))
	fmt.Printf("Add2(1,2)=%v\n", Add2(1,2))
	fmt.Printf("Add3(1,2)=%v\n", Add3(1,2))

	// 报错: multiple-value TupleAdd() in single-value context
	// fmt.Printf("TupleAdd(1,2,1)=%v\n", TupleAdd(1,2,1))
	// fmt.Printf("TupleAdd(1,2,1)=%v,%v\n", TupleAdd(1,2,1))

	var a,b = TupleAdd(1,2,1)
	fmt.Printf("TupleAdd(1,2,1)=%v,%v\n", a,b)
}