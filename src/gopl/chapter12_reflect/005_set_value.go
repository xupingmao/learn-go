package main


import (
	"fmt"
	"reflect"
	"strings"
)

func printValueInfo(value reflect.Value) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("value.kind = %v\n", value.Kind())
	fmt.Printf("value.type = %v\n", value.Type())
	fmt.Printf("value.interface = %t\n", value.Interface())
	// 是否可寻址
	fmt.Printf("value.CanAddr = %v\n", value.CanAddr())
}

func setObjectValue(value, newValue reflect.Value) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Set value by reflect.Value.Set")

	switch value.Kind() {
	case reflect.Ptr:
		fmt.Printf("value.Elem = %v\n", value.Elem())
		realValue := value.Elem()
		switch realValue.Kind() {
		case reflect.Int:
			// func (v reflect.Value) SetInt(x int)
			// func (v reflect.Value) SetString(x string)
			// ...
			realValue.Set(newValue)
			fmt.Printf("value = %v\n", value)
		default:
			fmt.Printf("unknown ptr\n")
		}
	default:
		fmt.Printf("unknown kind: %v\n", value.Kind())
		return
	}

	fmt.Printf("newValue = %v\n", newValue)
	fmt.Printf("value.Elem = %v\n", value.Elem())
}

func setValueByPtr(value, newValue reflect.Value) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Set value by pointer")

	switch value.Kind() {
	case reflect.Ptr:
		fmt.Printf("value.Elem = %v\n", value.Elem())
		realValue := value.Elem()
		switch realValue.Kind() {
		case reflect.Int:
			px := value.Interface().(*int)
			*px = newValue.Interface().(int)
		default:
			fmt.Printf("unknown ptr\n")
		}
	default:
		fmt.Printf("unknown kind: %v\n", value.Kind())
		return
	}

	fmt.Printf("newValue = %v\n", newValue)
	fmt.Printf("value.Elem = %v\n", value.Elem())
}

func main() {
	var x = 1

	printValueInfo(reflect.ValueOf(1))
	printValueInfo(reflect.ValueOf("hello"))
	printValueInfo(reflect.ValueOf(&x))

	setObjectValue(reflect.ValueOf(&x), reflect.ValueOf(10))
	setObjectValue(reflect.ValueOf(10), reflect.ValueOf(10))
}