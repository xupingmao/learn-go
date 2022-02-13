package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Movie struct {
	// JSON字段的名词必须大写，不然对外部而言是不可见的
	Title     string
	Year      int  `json:"released"`        // 这部分是tag，冒号(:)后面不能有空格，json库通过反射获取tag中的名称
	Color     bool `json:"color,omitempty"` // omitempty表示为0或者为空或者为false就不输出
	Actors    []string
	// 这个属性是小写字母开头的，是无法被反射读取的
	innerName string
}

var movies = []Movie{
	{
		Title:     "Casablanca",
		Year:      1942,
		Color:     false,
		Actors:    []string{"Humphrey Bogart", "Ingrid Bergman"},
		innerName: "innerName"},

	{
		Title:  "流浪地球",
		Year:   2019,
		Color:  true,
		Actors: []string{"屈楚萧", "赵今麦"},
	},
} // 注意go会自动给换行加分号(;)，通过添加逗号(,)可以避免插入分号


func printLineSep() {
	fmt.Println(strings.Repeat("-", 50))
}

func getInfoByReflection(x interface{}) {
	printLineSep()

	type stringer interface {
		String() string
	}

	var value = reflect.ValueOf(x)

	// 这里报错: syntax error: unexpected type, expecting name or (
	// fmt.Printf("Reflection x.type: %v\n", x.type)

	// 这里报错: use of .(type) outside type switch
	// fmt.Printf("Reflection x.type: %v\n", x.(type))

	fmt.Printf("Reflection x.type: %v\n", reflect.TypeOf(x)) 
	// 1) Reflection x.type: string
	// 2) Reflection x.type: []main.Movie

	// `reflect.TypeOf(x)` 返回 `*reflect.rtype` 类型的数据
	fmt.Printf("Reflection x.type.type: %v\n", reflect.TypeOf(reflect.TypeOf(x)))

	valueOfX := reflect.ValueOf(x) // 总是返回`reflect.Value`类型的数据
	fmt.Printf("Reflection x.value: %v\n", valueOfX)
	fmt.Printf("Reflection reflect.ValueOf(x).type: %v\n", reflect.TypeOf(valueOfX))

	// `reflect.Value.Interface` 是 `reflect.ValueOf` 的逆操作
	fmt.Printf("Reflection value.interface = %v\n", value.Interface())
	fmt.Printf("Reflection value.interface.(type) = %t\n", value.Interface())

	// var x2 reflect.Value = x
	// fmt.Printf("Reflection x2 = %t\n", x2)

	// `x.(type)` 是一个类型断言
	// `x.(int)` 类似于java的 `x instanceof int`
	switch x := x.(type) {
		case string:
			fmt.Printf("Reflection string: %v\n", x)
		case int:
			fmt.Printf("Reflection int: %v\n", x)
		default:
			fmt.Printf("Reflection unknown: %v\n", x)
	}
}

func getInfoByReflection2(x reflect.Value) {
	printLineSep()

	fmt.Printf("x.Kind = %v\n", x.Kind())
	fmt.Printf("x.Interface = %v\n", x.Interface())

	t := x.Type()

	switch x.Kind() {
	case reflect.Struct:
		for i:= 0; i < x.NumField(); i++ {
			fieldPath := fmt.Sprintf("(%s).%s", x.Type(), x.Type().Field(i).Name)
			fieldValue := x.Field(i)
			// StructTag定义
			// `type StructTag string`
			fieldTag := t.Field(i).Tag
			if fieldTag != "" {
				fmt.Printf("%s = %v `%s`\n", fieldPath, fieldValue, fieldTag)
			} else {
				fmt.Printf("%s = %v\n", fieldPath, fieldValue)
			}
		}
	}
}

func main() {
	// func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("json.data = %s\n", data)

	dataIndent, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("json.dataIndent = %s\n", dataIndent)

	getInfoByReflection("Hello,World")
	getInfoByReflection(movies)

	getInfoByReflection2(reflect.ValueOf("Hello,World"))
	getInfoByReflection2(reflect.ValueOf(movies))
	getInfoByReflection2(reflect.ValueOf(movies[0]))

	printLineSep()

	var movies2 []Movie

	// func Unmarshal(data []byte, v interface{}) error
	err = json.Unmarshal(data, &movies2)
	if err != nil {
		fmt.Printf("JSON umarshaling failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("obj.movies2 = %v\n", movies2)
}
