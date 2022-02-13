package main

import (
	"fmt"
	"strings"
	"os"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1,j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArray() {
	fmt.Println(strings.Repeat("-", 50))
	var arr = [5]int{3, 5, 2, 10, 1}

	fmt.Printf("[1] arr = %v\n", arr)
	
	// slice可以改源数组的值
	reverse(arr[:])

	fmt.Printf("[2] arr = %v\n", arr)
}

func testArray() {
	fmt.Println(strings.Repeat("-", 50))
	var arr = [...]int{3,5,2}
	fmt.Printf("arr = %t(%T)\n", arr, arr)
}

func testSlice() {
	fmt.Println(strings.Repeat("-", 50))
	var arr []string
	var arr2 = append(arr, "a") // 这是一个新的对象

	fmt.Printf("arr = %t(%T), addr:%p\n", arr, arr, arr)
	fmt.Printf("arr2 = %t, addr:%p\n", arr2, arr2)

	// 判断 arr == arr2 报错
	// (slice can only be compared to nil)

	if &arr == &arr2 {
		fmt.Fprintf(os.Stderr, "arr can not equal arr2")
		os.Exit(1)
	}
}

func testSliceModify() {
	// slice本身不会存储数据，而是引用数组
	// slice有三个属性 {data *[]type, len int; cap int;}
	fmt.Println(strings.Repeat("-", 50))
	var arr1 = []int{1,2}
	var arr2 = arr1[:]

	fmt.Printf("arr2 = %v\n", arr2)
	arr1[0] = 2 // 修改之后 arr2[0] 也变成了 2
	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)
}

func testSliceCap() {
	fmt.Println(strings.Repeat("-", 50))
	var array = [...]int{1,2,3,4,5,6,7,8,9,0}

	var slice1 = array[0:3]

	fmt.Printf("array = %v\n", array)
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("cap(slice1) = %v\n", cap(slice1))

	var slice2 = append(slice1, 10)
	fmt.Printf(">>> var slice2 = append(slice1, 10)\n")
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)
	fmt.Printf("array[0:4] = %v\n", array[0:4])
	fmt.Printf("array = %v\n", array)
	fmt.Printf("array[3] has been modified\n")
	fmt.Printf("&array = %v\n", &array)
	var slice3 = array[:]
	fmt.Printf(">>> var slice3 = array[:]\n")
	fmt.Printf("&slice3 = %p\n", slice3)
	fmt.Printf(">>> var slice4 = append(slice3, 11)\n")
	var slice4 = append(slice3, 11)
	fmt.Printf("slice4.addr = %p, &slice4 = %p\n", slice4, &slice4)
	fmt.Printf("slice4.addr = %p, &slice4 = %p\n", slice4, &slice4)
}

func main() {
	testArray()

	reverseArray()

	testSlice()

	testSliceModify()

	testSliceCap()
}
