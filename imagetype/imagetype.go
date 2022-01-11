package imagetype

import (
	"errors"
	"image/color"
)

/**
 * RGB
 * 0~255までの値を入れRBGを表現
 */
type Rgb struct {
	Red   uint8
	Green uint8
	Blue  uint8
	A     uint8
}

/**
 * Materialtype
 *
 * 追加する素材はここに追加する
 * 0 = スケルトン
 * 1 = マット
 */

/**
 * Position
 * オブジェクトの中心位置
 */
type Position struct {
	X float32
	Y float32
	Z float32
}

/**
 * Square
 * 立方体の値を入れる
 */
type Square struct {
	Objectname string
	Pos        Position
	AngleX     float32
	AngleY     float32
	AngleZ     float32
	Color      Rgb
	Length     float32
	Material   int
}

/**
 * Sphia
 * 球体の値を入れる
 */
type Sphia struct {
	Objectname string
	Pos        Position
	Color      Rgb
	Length     float32
	Material   int
}

/**
 * Light
 * 光の値を入れる
 */
type Light struct {
	Objectname string
	Pos        Position
	Color      Rgb
	Bright     float64
}

/**
 * Light
 * 床の情報
 * direction = x->y->z
 */
type Floor struct {
	Direction int
	Pos       float32
	Color     Rgb
	Material  int
}

/**
 * Camera
 * カメラの情報
 */
type Camera struct {
	Pos    Position
	AngleX int
	AngleY int
	AngleZ int
}

/*--------------------------------------*
 *   以下オブジェクトのリスト定義         *
 *--------------------------------------*/

var fl Floor
var li Light
var sq []Square
var Cm Camera
var sphia []Sphia

//s=sleton,m=mat
var RenderType string

// 出力ファイルの名前
var Filename string

// 背景色
var Background Rgb

// 横
var W int = 400

// 縦
var H int = 400

func SetCamera() bool {
	Cm.AngleX = 0
	Cm.AngleY = 0
	Cm.AngleZ = 0
	Cm.Pos.X = 0
	Cm.Pos.Y = 1.4
	Cm.Pos.Z = 0
	return true
}

// 床のセッター
func SetFloor(f Floor) bool {
	fl = f
	return true
}

// 光のセッター
func SetLight(l Light) bool {
	li = l
	return true
}

// 立方体のセッター
func SetSquare(s Square) bool {
	sq = append(sq, s)
	return true
}

// 球体のセッター
func SetSphia(s Sphia) bool {
	sphia = append(sphia, s)
	return true
}

// 球体のゲッター
func GetSphia(x int) (Sphia, error) {
	if len(sphia) > x {
		return sphia[x], nil
	}
	return sphia[0], errors.New("配列の外です")
}

func GetSphiaLen() int {
	return len(sphia)
}

func GetSphias() []Sphia {
	return sphia
}

// 床のゲッター
func GetFloor() Floor {
	return fl
}

// 光のゲッター
func GetLight() Light {
	return li
}

// 立方体のゲッター
func GetSquare(zz int) (Square, error) {
	if len(sq) > zz {
		return sq[zz], nil
	}
	return sq[0], errors.New("配列の範囲外です")

}

func GetLenSquare() int {
	return len(sq)
}

// ポジションのセッター
func SetPos(x, y, z float32) Position {
	var p Position
	p.X = x
	p.Y = y
	p.Z = z
	return p
}

// ポジションのゲッター
func GetPos(p Position) []float32 {
	var a []float32
	a = append(a, p.X)
	a = append(a, p.Y)
	a = append(a, p.Z)
	return a
}

// colorすべてを一括で設定するセッター
func SetColorAll(a byte) Rgb {
	var n Rgb
	n.Blue = a
	n.Green = a
	n.Red = a
	n.A = 0xff
	return n
}

// Rを設定するセッター
func SetColorR(a byte, n Rgb) Rgb {
	n.Red = a
	return n
}

// Gを設定するセッター
func SetColorG(a byte, n Rgb) Rgb {
	n.Green = a
	return n
}

// Bを設定するセッター
func SetColorB(a byte, n Rgb) Rgb {
	n.Blue = a
	return n
}

func GetChangeRGBA(n Rgb) color.RGBA {
	var c color.RGBA
	c.R = n.Red
	c.G = n.Green
	c.B = n.Blue
	c.A = n.A
	return c
}
