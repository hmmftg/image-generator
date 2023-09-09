package imagegenerator

import (
	"image"

	"golang.org/x/image/font"
)

type Drawable interface {
	Draw(tx *PrintTx) int
	Font() string
}

type PrintRequest struct {
	Drawings []Drawable
	Images   []ImageData
}

type PrintTx struct {
	Rgba  *image.RGBA
	Src   image.Image
	Fg    image.Image
	Bg    *image.Uniform
	Fonts *map[string]font.Face
}
