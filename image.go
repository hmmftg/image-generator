package imagegenerator

import (
	"image"

	"log"

	"github.com/hmmftg/image/draw"
)

type Image struct {
	ID         string
	Target     image.Image
	Scale      float64
	X, Y       float64
	RightAlign bool
}

func DrawImage(img draw.Image, target image.Image, x, y int) {
	r := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + target.Bounds().Dx(), Y: y + target.Bounds().Dy()}}
	draw.Draw(img, r, target, image.Point{0, 0}, draw.Over)
}

func (i Image) Draw(tx *PrintTx) int {
	if i.Scale == 0 {
		i.Scale = 1
	}
	var ok bool
	i.Target, ok = tx.Images[i.ID]
	if !ok {
		log.Println("image not found(ignored)", i.ID)
		return 0
	}
	scaledRect := image.Rect(
		0,
		0,
		int(float64(i.Target.Bounds().Dx()*int(tx.Dpi/72.))*i.Scale),
		int(float64(i.Target.Bounds().Dy()*int(tx.Dpi/72.))*i.Scale))
	// log.Println("scaledRect", scaledRect, tx.Dpi/72., i.Scale, math.Ceil((tx.Dpi/72.)*i.Scale))
	scaledImage := tx.getNewImage(scaledRect)

	draw.NearestNeighbor.Scale(scaledImage, scaledRect, i.Target, i.Target.Bounds(), draw.Over, nil)

	x := tx.RelationalX(i.X)
	if i.RightAlign {
		x = tx.Rgba.Bounds().Max.X - scaledImage.Bounds().Max.X - x
	} else {
		x = tx.RelationalX(i.X)
	}
	y := tx.RelationalY(i.Y)

	DrawImage(tx.Rgba, scaledImage, x, y)
	return x
}
