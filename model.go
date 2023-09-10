package imagegenerator

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Drawable interface {
	Draw(tx *PrintTx) int
}

type PrintRequest struct {
	Drawings []Drawable
	Images   []ImageData
}

type PrintTx struct {
	Rgba   *image.RGBA
	Src    image.Image
	Fg     image.Image
	Bg     *image.Uniform
	Dpi    float64
	Fonts  map[string]opentype.Font
	Faces  *map[string]font.Face
	Images map[string]image.Image
}
