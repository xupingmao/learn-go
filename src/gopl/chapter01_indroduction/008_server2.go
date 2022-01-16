package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// 问题：浏览器请求/count的时候，两个处理器都会被触发，但是只输出了counter的结果
	// 在浏览器请求/count的时候，浏览器会同时请求/favion.ico，导致触发了handler处理器
	// 通过curl命令来触发的时候，请求/count只会触发 counter 处理器
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Printf("%05d: handler.start\n", count)
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "Count = %d\n", count)

	fmt.Printf("%05d: handler.end\n", count)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Printf("%05d: counter.start\n", count)
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	fmt.Printf("%05d: counter.end\n", count)
}


