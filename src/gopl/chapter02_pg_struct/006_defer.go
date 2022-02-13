package main

import "fmt"

func closeResource1(value int) {
    fmt.Printf("defer close1 i=%d\n", value)
}

func closeResource2(ptr *int) {
    fmt.Printf("defer close2 ptr i=%d\n", *ptr)
}

func main() {
    i:=0
    // defer 有点像C++的析构函数
    // defer 后的表达式的参数是声明的时候的值，可以理解为把后面表达式创建成了一个任务，
    // 在退出函数的时候执行，而且执行的顺序是和声明顺序反向的
    defer closeResource1(i)
    defer closeResource2(&i)

    i += 1
    fmt.Printf("i=%d\n", i)
}

// 输出
// i=1
// defer close2 ptr i=1
// defer close1 i=0