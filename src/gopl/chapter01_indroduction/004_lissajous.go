package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Go语言标准库参考手册
// https://studygolang.com/pkgdoc

// color.RGBA{R, G, B, Alpha}
var green = color.RGBA{0, 0xff, 0, 0xff}
var red = color.RGBA{0xff, 0, 0, 0xff}
var blue = color.RGBA{0, 0, 0xff, 0xff}
var palette = []color.Color{color.White, color.Black, green, red, blue}

const (
	whiteIndex = 0 // 画板的第一种颜色
	blackIndex = 1 // 画板的第二种颜色
	greenIndex = 2 // 绿色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	color := "green"
	if len(os.Args) > 2 {
		color = os.Args[2]
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(strings.Repeat("-", 50))
			fmt.Printf("request received: %v\n", r)
			lissajous(w, color)
		}
		http.HandleFunc("/", handler)

		fmt.Println("start server...")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}

	lissajous(os.Stdout, color)
}

func lissajous(out io.Writer, color string) {
	const (
		cycles  = 5     // 完整的x震荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 画像画布包含[-size..+size]
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 阶段的差异
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		if color != "black" {
			// 绘制背景
			for x := 0; x < size*2+1; x++ {
				for y := 0; y < size*2+1; y++ {
					img.SetColorIndex(x, y, blackIndex)
				}
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// 原版白底黑色
			if color == "black" {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					blackIndex)
			}

			if color == "green" {
				// 绿色
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					greenIndex)
			}

			if color == "random" {
				randColorIndex := uint8(rand.Int() % len(palette))
				// 随机颜色
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					randColorIndex)
			}

		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // 这里忽略了编码错误

}
