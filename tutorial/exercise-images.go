package main

import "golang.org/x/tour/pic"
import (
	"image"
	"image/color"
)

// Imageに必要な要素は大きさだけ
type Image struct{
	width, height int
}

// これがないと，ShowImageにかからない
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// これがないと，ShowImageにかからない
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

// メモリで持っている必要はなくて，場所を聞かれたらその場所を返せば良い．
// 他のやり方はなくね？事前にどこかにでかい配列を作って，そこ経由で飛ばすこともできなくはないけど，やるにしてもどこ？という気持ちがある．
func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}


func main() {
	m := Image{200, 300}
	pic.ShowImage(m)
}
