package imagegenerator

import (
	"image"
	"image/draw"
	"log"
)

type Image struct {
	ID     string
	Target image.Image
	X, Y   int
}

func DrawImage(img *image.RGBA, target image.Image, x, y int) {
	r := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + target.Bounds().Dx(), Y: y + target.Bounds().Dy()}}
	draw.Draw(img, r, target, image.Point{0, 0}, draw.Over)
}

func (i Image) Draw(tx *PrintTx) int {
	var ok bool
	i.Target, ok = tx.Images[i.ID]
	if !ok {
		log.Println("image not found", i.ID)
		return 0
	}
	DrawImage(tx.Rgba, i.Target, i.X, i.Y)
	return i.X
}
