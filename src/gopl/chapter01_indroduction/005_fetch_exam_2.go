package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func formatUrl(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}

func main() {
	for _, url := range os.Args[1:] {
		url = formatUrl(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		nbytes, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("nbytes:%v url:%s\n", nbytes, url)
	}
}

// $> go run 005_fetch_exam_2.go baidu.com
// <html>
// <meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
// </html>
// nbytes:81 url:http://baidu.com

