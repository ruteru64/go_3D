package main

import (
	"go3D/imagetype"
	"go3D/render"
	"os"
)

/**
 * [x,y,z]0,1,0の位置にカメラを置き
 * \Θ| |Θ/ 上下同様
 * Θ = 30°のなるように平行にカメラを置く(後々修整予定)
 * コマンドライン引数に"clean"がある場合生成したファイルを削除
 */
func main() {
	test()
	render.Do()
	if len(os.Args) == 2 {
		if os.Args[1] == "clean" {
			os.Remove(imagetype.Filename)
		}
	}
	return
}

func test() {
	imagetype.Filename = "test.png"
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
	c.Length = 20.0
	c.AngleX = 30
	c.AngleY = 320
	c.AngleZ = 70
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
}
