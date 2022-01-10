package render

import (
	"errors"
	"fmt"
	"go3D/imagetype"
	"image"
	"image/png"
	"math"
	"os"
)

var err error = nil

func Do() {
	var a image.Rectangle
	a.Max = image.Point{imagetype.W, imagetype.H}
	a.Min = image.Point{0, 0}
	myImage := image.NewRGBA(a)
	i := 0
	for k := 0; k < len(myImage.Pix); k += 4 {
		i++
		if k%(400*4) == 0 {
			i = 0
		}
		myImage.Pix[k] = imagetype.Background.Red
		myImage.Pix[k+1] = imagetype.Background.Green
		myImage.Pix[k+2] = imagetype.Background.Blue
		myImage.Pix[k+3] = imagetype.Background.A
	}
	var w [511][511][511]int8 // [x][y][z] 0.1が最小単位
	d := createCube(w)
	if err != nil {
		return
	}
	Xcode(d, myImage)
	savefile, err := os.Create(imagetype.Filename)
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer savefile.Close()
	png.Encode(savefile, myImage)
	return
}

func Xcode(d [511][511][511]int8, n *image.RGBA) bool {
	f := -1
	for i := 55; i < 400+55; i++ {
		for k := 55; k < 400+55; k++ {
			f++
			for m := 55; m < 400+55; m++ {
				if d[400+55-k][400+55-i][m] == 2 {
					n.Pix[f*4] = 0
					n.Pix[f*4+1] = 0
					n.Pix[f*4+2] = 0
					//fmt.Println(k, i)
				} else if d[400+55-k][400+55-i][m] == 3 && !(n.Pix[f*4] == 0 && n.Pix[f*4+1] == 0 && n.Pix[f*4+2] == 0) {
					n.Pix[f*4] = 0 + 0xff
					n.Pix[f*4+1] = 0 + 0xff
					n.Pix[f*4+2] = 0 + 0xff
				}
			}
			if int(imagetype.GetFloor().Pos) >= 400+55-i-1 {
				n.Pix[f*4] = imagetype.GetFloor().Color.Red
				n.Pix[f*4+1] = imagetype.GetFloor().Color.Green
				n.Pix[f*4+2] = imagetype.GetFloor().Color.Blue
			}
		}
	}
	return true
}

func createCube(d [511][511][511]int8) [511][511][511]int8 {
	for index := 0; index < imagetype.GetLenSquare(); index++ {
		sq := imagetype.GetSquare(index)
		size := sq.Length
		marginX := int((sq.Pos.X - size/2 + 5.6) * 10)
		marginZ := int((sq.Pos.Z - size/2 + 5.6) * 10)
		marginY := int((sq.Pos.Y - size/2 + 5.6) * 10)
		if marginX >= 0 {
			marginX = 0
		} else {
			marginX *= -1
		}
		if marginZ >= 0 {
			marginZ = 0
		} else {
			marginZ *= -1
		}
		if marginY >= 0 {
			marginY = 0
		} else {
			marginY *= -1
		}
		var poss [3 * 8]int
		for i := 0; i < 8; i++ {
			var pos imagetype.Position
			/*
				//テストコード
				if i != 6 {
					if true {
						continue
					}
				}
			*/
			if i/4 == 0 { //右
				if i/2 == 0 { //前
					if i%2 == 0 { //上
						pos.X = float32(math.Cos((float64(sq.AngleY)+45)/180)*math.Cos((float64(sq.AngleZ)-45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)-45)/180)*math.Cos((float64(sq.AngleX)+45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)+45)/180)*math.Cos((float64(sq.AngleY)+45)/180)) * size
					} else { //下
						pos.X = float32(math.Cos((float64(sq.AngleY)-45)/180)*math.Cos((float64(sq.AngleZ)-45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)-45)/180)*math.Cos((float64(sq.AngleX)+45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)+45)/180)*math.Cos((float64(sq.AngleY)-45)/180)) * size
					}
				} else { //後
					if i%2 == 0 { //上
						pos.X = float32(math.Cos((float64(sq.AngleY)+45)/180)*math.Cos((float64(sq.AngleZ)+45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)+45)/180)*math.Cos((float64(sq.AngleX)+45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)+45)/180)*math.Cos((float64(sq.AngleY)+45)/180)) * size
					} else { //下
						pos.X = float32(math.Cos((float64(sq.AngleY)-45)/180)*math.Cos((float64(sq.AngleZ)+45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)+45)/180)*math.Cos((float64(sq.AngleX)+45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)+45)/180)*math.Cos((float64(sq.AngleY)-45)/180)) * size
					}
				}
			} else { //左
				if i/2 == 3 { //前
					if i%2 == 0 { //上
						pos.X = float32(math.Cos((float64(sq.AngleY)+45)/180)*math.Cos((float64(sq.AngleZ)-45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)-45)/180)*math.Cos((float64(sq.AngleX)-45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)-45)/180)*math.Cos((float64(sq.AngleY)+45)/180)) * size
					} else { //下
						pos.X = float32(math.Cos((float64(sq.AngleY)-45)/180)*math.Cos((float64(sq.AngleZ)-45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)-45)/180)*math.Cos((float64(sq.AngleX)-45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)-45)/180)*math.Cos((float64(sq.AngleY)-45)/180)) * size
					}
				} else { //後
					if i%2 == 0 { //上
						pos.X = float32(math.Cos((float64(sq.AngleY)+45)/180)*math.Cos((float64(sq.AngleZ)+45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)+45)/180)*math.Cos((float64(sq.AngleX)-45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)-45)/180)*math.Cos((float64(sq.AngleY)+45)/180)) * size
					} else { //下
						pos.X = float32(math.Cos((float64(sq.AngleY)-45)/180)*math.Cos((float64(sq.AngleZ)+45)/180)) * size
						pos.Y = float32(math.Cos((float64(sq.AngleZ)+45)/180)*math.Cos((float64(sq.AngleX)-45)/180)) * size
						pos.Z = float32(math.Cos((float64(sq.AngleX)-45)/180)*math.Cos((float64(sq.AngleY)-45)/180)) * size
					}
				}
			}
			poss[i*3] = int((pos.X+sq.Pos.X)*10) + 56
			poss[i*3+1] = int((pos.Y+sq.Pos.Y)*10) + 56
			poss[i*3+2] = int((pos.Z+sq.Pos.Z)*10) + 56
			if poss[i*3] < 0 || poss[i*3+1] < 0 || poss[i*3+2] < 0 || poss[i*3] >= 511 || poss[i*3+1] >= 511 || poss[i*3+2] >= 511 {
				fmt.Println("値の範囲オーバー")
				err = errors.New("値の範囲オーバー")
				return d
			}
			d[poss[i*3]][poss[i*3+1]][poss[i*3+2]] = 2
			//fmt.Println(poss)
		}
		var connect [7][3]int
		connect[0] = [3]int{1, 2, 6}
		connect[1] = [3]int{3, 7, 0}
		connect[2] = [3]int{3, 4, 0}
		connect[3] = [3]int{5, 0, 0}
		connect[4] = [3]int{5, 6, 0}
		connect[5] = [3]int{7, 0, 0}
		connect[6] = [3]int{7, 0, 0}
		for i := 0; i <= 6; i++ {
			for f := 0; f < 3; f++ {
				k := connect[i][f]
				if k == 0 {
					continue
				}
				x := poss[i*3] - poss[k*3]
				y := poss[i*3+1] - poss[k*3+1]
				z := poss[i*3+2] - poss[k*3+2]
				xa := int(size) * 4
				for n := 1; n < xa; n++ {
					d[poss[i*3]-(x*n)/xa][poss[i*3+1]-(y*n)/xa][poss[i*3+2]-(z*n)/xa] = 2
				}
			}
		}
		if sq.Material == 1 {
			//fmt.Println("中を埋める")
			surface := [24]int{0, 1, 2, 3, 0, 1, 6, 7, 0, 2, 6, 4, 1, 3, 7, 5, 2, 3, 4, 5, 4, 5, 6, 7}
			for i := 0; i < 6; i++ {
				x := poss[surface[i*4]*3] - poss[surface[i*4+1]*3]
				y := poss[surface[i*4]*3+1] - poss[surface[i*4+1]*3+1]
				z := poss[surface[i*4]*3+2] - poss[surface[i*4+1]*3+2]
				x1 := poss[surface[i*4]*3] - poss[surface[i*4+2]*3]
				y1 := poss[surface[i*4]*3+1] - poss[surface[i*4+2]*3+1]
				z1 := poss[surface[i*4]*3+2] - poss[surface[i*4+2]*3+2]
				xa := int(size) * 4
				for i0 := 1; i0 < int(size)*4-1; i0++ {
					for i1 := 1; i1 < int(size)*4-1; i1++ {
						if d[poss[surface[i*4]*3]-(x*i0)/xa-(x1*i1)/xa][poss[surface[i*4]*3+1]-(y*i0)/xa-(y1*i1)/xa][poss[surface[i*4]*3+2]-(z*i0)/xa-(z1*i1)/xa] != 2 {
							d[poss[surface[i*4]*3]-(x*i0)/xa-(x1*i1)/xa][poss[surface[i*4]*3+1]-(y*i0)/xa-(y1*i1)/xa][poss[surface[i*4]*3+2]-(z*i0)/xa-(z1*i1)/xa] = 3
						}
					}
				}
			}
		}
	}
	return d
}
