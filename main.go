package main

import (
	"fmt"
	"go3D/imagetype"
	"go3D/input"
	"go3D/render"
	"net/http"
	"os"
	"time"
)

// httpに適応した,適当に増やす
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/genarate", handlerGenarate)
	http.ListenAndServe(":80", nil)
}

// / の時
func handler(web http.ResponseWriter, request *http.Request) {
	fmt.Println("Time:" + time.Now().String() + " endpoint:\"/\"")
	fmt.Fprintf(web, "404 NOT FOUND")
}

// /index の時
func handlerIndex(web http.ResponseWriter, request *http.Request) {
	fmt.Println("Time:" + time.Now().String() + " endpoint:\"/index\"")
	fp, err := os.Open("./public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	s := make([]byte, 1024)
	fp.Read(s)
	fmt.Fprintf(web, string(s))
}

// /genarate の時
func handlerGenarate(web http.ResponseWriter, request *http.Request) {
	fmt.Println("Time:" + time.Now().String() + " endpoint:\"/genarate\"")
	generate()
	fp, err := os.Open("./test.png")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	img := make([]byte, 1024*1024*1024)
	fp.Read(img)
	web.Write(img)
}

/**
 * [x,y,z]0,1,0の位置にカメラを置き
 * \Θ| |Θ/ 上下同様
 * Θ = 30°のなるように平行にカメラを置く(後々修整予定)
 * コマンドライン引数に"clean"がある場合生成したファイルを削除
 */
func generate() {
	if input.IsGenalateThisFile {
		input.SetInput()
	} else {
		test()
	}
	render.Do()
	return
}

func clean() {
	os.Remove(imagetype.Filename)
}

// 画像生成のテストコードinput.SetInput()で試用不使用を指定できる
func test() {
	imagetype.Filename = "test.png"
	imagetype.W = 800
	imagetype.H = 800
	var v imagetype.Rgb
	imagetype.RenderType = "s"
	v = imagetype.SetColorAll(0x88)
	v = imagetype.SetColorR(107, v)
	v = imagetype.SetColorG(216, v)
	v = imagetype.SetColorB(255, v)
	imagetype.Background = v
	var a imagetype.Floor
	a.Direction = 0
	a.Color = imagetype.SetColorAll(0x88)
	a.Material = 1
	a.Pos = 10.0
	imagetype.SetFloor(a)

	var b imagetype.Light
	b.Color = imagetype.SetColorAll(0xff)
	b.Bright = 10.0
	b.Objectname = "testlight"
	b.Pos = imagetype.SetPos(10.0, 10.0, 10.0)
	imagetype.SetLight(b)

	var c imagetype.Square
	c.Pos = imagetype.SetPos(10.0, 10.0, 20.0)
	c.Length = 50.0
	c.AngleX = 30
	c.AngleY = 32
	c.AngleZ = 30
	c.Material = 0
	c.Color = imagetype.SetColorAll(0xff)
	c.Objectname = "testCube"
	imagetype.SetSquare(c)
	c.Pos = imagetype.SetPos(10.0, 10.0, 20.0)
	c.Length = 10.0
	c.AngleX = 50
	c.AngleY = 30
	c.AngleZ = 20
	c.Material = 1
	c.Color = imagetype.SetColorAll(0xff)
	c.Objectname = "testCube2"
	imagetype.SetSquare(c)
	var s imagetype.Sphia
	s.Pos = imagetype.SetPos(5.0, 5.0, 0.0)
	s.Color = imagetype.SetColorAll(0xff)
	s.Length = 5.0
	s.Material = 1
	s.Objectname = "testSphia"
	imagetype.SetSphia(s)
}
