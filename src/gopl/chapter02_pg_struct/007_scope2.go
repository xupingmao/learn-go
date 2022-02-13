package main

import (
	"os"
	"log"
)

var cwd string
var ccc string

func f() {	
	// 这里的cwd是新生成的变量
    cwd, err := os.Getwd()
    if err != nil {
    	log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory: %v", cwd)
}

func f2() {
	// 如果想要赋值给全局变量cwd应该这样使用
	var err error
    cwd, err = os.Getwd()
    if err != nil {
    	log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory: %v", cwd)
}

func main() {
	f()
	log.Printf("global.cwd: %v", cwd)
	f2()
	log.Printf("global.cwd: %v", cwd)
	// 这个打印出来的是空字符串，但是不打印也不报错
	// 全局变量不报`declared and not used`错误
}

// 输出
// 2022/01/26 21:56:48 Working directory: /Users/mark/projects/learn-go