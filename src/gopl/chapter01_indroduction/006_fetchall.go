package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)


func formatUrl(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}

// chan string 和 chan <- string的区别是啥？
// 直接使用chan string也是可以执行的
// chan string是双向管道
// chan <- string 是只发送管道 send-only
func fetch(url string, ch chan <- string) {
	start := time.Now()
	url = formatUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


func main() {
	start := time.Now()
	// chan是Go语言的关键字，代表管道
	// make是函数，用于创建对象
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// 正常情况是需要关闭管道的
	// 这里按照参数的数量来读取管道，和生成的go程一致，实际上管道并未关闭
	for range os.Args[1:] {
		// 从管道读取数据是阻塞的
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// 正常情况
// $> go run 006_fetchall.go https://baidu.com https://qq.com
// 0.44s  132549 https://qq.com
// 1.54s  328669 https://baidu.com
// 1.54s elapsed

// 练习1.11
// $> go run 006_fetchall.go http://noshu.sh
// Get http://noshu.sh: dial tcp: lookup noshu.sh: no such host
// 0.85s elapsed
