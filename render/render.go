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

var d [1000][1000][1000]int8 // [x][y][z] 0.1が最小単位

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
	err := createCube()
	if err != nil {
		return
	}
	err = createSphia()
	if err != nil {
		return
	}
	Xcode(myImage)
	savefile, err := os.Create(imagetype.Filename)
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer savefile.Close()
	png.Encode(savefile, myImage)
	return
}

func Xcode(n *image.RGBA) bool {
	f := -1
	if imagetype.W > 900 && imagetype.H > 900 {
		fmt.Println("画像サイズが大きすぎます")
		return false
	}
	for i := 100; i < imagetype.H+100; i++ {
		for k := 100; k < imagetype.W+100; k++ {
			f++
			for m := 100; m < imagetype.W+100; m++ {
				if d[imagetype.W+100-k][imagetype.H+100-i][m] == 2 {
					n.Pix[f*4] = 0
					n.Pix[f*4+1] = 0
					n.Pix[f*4+2] = 0
					//fmt.Println(k, i)
				} else if d[imagetype.H+100-k][imagetype.W+100-i][m] == 3 && !(n.Pix[f*4] == 0 && n.Pix[f*4+1] == 0 && n.Pix[f*4+2] == 0) {
					n.Pix[f*4] = 0 + 0xff
					n.Pix[f*4+1] = 0 + 0xff
					n.Pix[f*4+2] = 0 + 0xff
				}
			}
			if int(imagetype.GetFloor().Pos) >= imagetype.H+99-i {
				n.Pix[f*4] = imagetype.GetFloor().Color.Red
				n.Pix[f*4+1] = imagetype.GetFloor().Color.Green
				n.Pix[f*4+2] = imagetype.GetFloor().Color.Blue
			}
		}
	}
	return true
}

func createCube() error {
	for index := 0; index < imagetype.GetLenSquare(); index++ {
		sq, err := imagetype.GetSquare(index)
		if err != nil {
			fmt.Println(err)
			return err
		}
		size := sq.Length
		marginX := int((sq.Pos.X - size/2 + 10) * 10)
		marginZ := int((sq.Pos.Z - size/2 + 10) * 10)
		marginY := int((sq.Pos.Y - size/2 + 10) * 10)
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
			poss[i*3] = int((pos.X+sq.Pos.X)*10) + 100
			poss[i*3+1] = int((pos.Y+sq.Pos.Y)*10) + 100
			poss[i*3+2] = int((pos.Z+sq.Pos.Z)*10) + 100
			if poss[i*3] < 0 || poss[i*3+1] < 0 || poss[i*3+2] < 0 || poss[i*3] >= 1000 || poss[i*3+1] >= 1000 || poss[i*3+2] >= 1000 {
				fmt.Println("値の範囲オーバー")
				return errors.New("値の範囲オーバー")
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
	return nil
}

// 球体を配列に入れる処理を記述
func createSphia() error {
	for index := 0; index < imagetype.GetSphiaLen(); index++ {
		sph, err := imagetype.GetSphia(index)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		size := abs(sph.Length)
		if (sph.Pos.X+size)*10 > 1000 || (sph.Pos.Y+size)*10 > 1000 || (sph.Pos.Z+size)*10 > 1000 {
			err = errors.New("範囲外です")
			fmt.Println(err)
			return err
		}
		if (sph.Pos.X-size)*10+100 < 0 || (sph.Pos.Y-size)*10+100 < 0 || (sph.Pos.Z-size)*10+100 < 0 {
			err = errors.New("範囲外です")
			fmt.Println(err)
			return err
		}
		if sph.Material == 0 {
			d[int(sph.Pos.X*10)+100][int(sph.Pos.Y*10)+100][int(sph.Pos.Z*10)+100] = 2
			d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
			d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
			d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
			d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
			d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
			d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
			d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
			d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
			for i := 0; i < int(size)*5; i++ {
				d[int((sph.Pos.X+size)*10)+100-i*4][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100-i*4][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100-i*4][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z+size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100-i*4][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z-size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y+size)*10)+100-i*4][int((sph.Pos.Z+size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y+size)*10)+100-i*4][int((sph.Pos.Z-size)*10)+100] = 2
				d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y+size)*10)+100-i*4][int((sph.Pos.Z+size)*10)+100] = 2
				d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y+size)*10)+100-i*4][int((sph.Pos.Z-size)*10)+100] = 2
				d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z+size)*10)+100-i*4] = 2
				d[int((sph.Pos.X+size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z+size)*10)+100-i*4] = 2
				d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y+size)*10)+100][int((sph.Pos.Z+size)*10)+100-i*4] = 2
				d[int((sph.Pos.X-size)*10)+100][int((sph.Pos.Y-size)*10)+100][int((sph.Pos.Z+size)*10)+100-i*4] = 2
			}
			for i := 0; i < 120; i++ {
				d[int((sph.Pos.X+float32(math.Sin(float64(i*3)))*size)*10)+100][int(sph.Pos.Y*10)+100][int((sph.Pos.Z+float32(math.Cos(float64(i*3)))*size)*10)+100] = 2
				d[int((sph.Pos.X+float32(math.Cos(float64(i*3)))*size)*10)+100][int((sph.Pos.Y+float32(math.Sin(float64(i*3)))*size)*10)+100][int(sph.Pos.Z*10)+100] = 2
				d[int((sph.Pos.X)*10)+100][int((sph.Pos.Y+float32(math.Cos(float64(i*3)))*size)*10)+100][int((sph.Pos.Z+float32(math.Sin(float64(i*3)))*size)*10)+100] = 2
			}
		} else {
			for i := 0; i < 360; i++ {
				for i0 := 0; i0 < 360; i0++ {
					for i1 := 0; i1 < 360; i1++ {
						d[int((sph.Pos.X+float32(math.Sin(float64(i)))*float32(math.Cos(float64(i0)))*size)*10)+100][int((sph.Pos.Y+float32(math.Sin(float64(i0)))*float32(math.Cos(float64(i1)))*size)*10)+100][int((sph.Pos.Z+float32(math.Sin(float64(i1)))*float32(math.Cos(float64(i)))*size)*10)+100] = 2
					}
				}
			}
		}
	}
	return nil
}

func abs(c float32) float32 {
	if c >= 0 {
		return c
	}
	return -1 * c
}
