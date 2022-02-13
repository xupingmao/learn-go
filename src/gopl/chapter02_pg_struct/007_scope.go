package main

import "fmt"

var g int = 10

func f() {
    // 寻找变量时会按照由内到外的顺序查找
    // 这里先查找f函数的临时变量，然后寻找全局变量
    var g int = 20
    fmt.Printf("func f(): g=%v\n", g)
}

func main() {
    f()
    fmt.Printf("func main(): g=%v\n", g)
}

// 输出
// func f(): g=20
// func main(): g=10