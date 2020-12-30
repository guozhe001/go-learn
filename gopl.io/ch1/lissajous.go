// 运行方式1: go run lissajous.go > out.gif
// 运行方式2: go run lissajous.go web
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
	"time"
)

// 12页 练习1.5 改变颜色; 12页 练习1.5添加更多颜色
var palette = []color.Color{color.RGBA{0x11, 0xff, 0x11, 0xff}, color.Black, color.RGBA{0xfb, 0x23, 0xdf, 0xff}}

// var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		fmt.Println("please open: http://localhost:8000")
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   //图像画布包含[-size..+size]
		nframes = 64    //动画中的帧数
		delay   = 8     //以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//  12页 练习1.5添加更多颜色，修改颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
