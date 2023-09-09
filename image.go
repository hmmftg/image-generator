package imagegenerator

import (
	"image"
	"image/draw"
)

type Image struct {
	Img    *image.RGBA
	Target image.Image
	X, Y   int
}

func DrawImage(img *image.RGBA, target image.Image, x, y int) {
	r := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + target.Bounds().Dx(), Y: y + target.Bounds().Dy()}}
	draw.Draw(img, r, target, image.Point{0, 0}, draw.Over)
}

func (i Image) Draw() {
	DrawImage(i.Img, i.Target, i.X, i.Y)
}
