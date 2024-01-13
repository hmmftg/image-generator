package imagegenerator

import (
	"image"
	"image/color"

	"github.com/hmmftg/image/draw"
	"github.com/hmmftg/image/font"
	"github.com/hmmftg/image/font/opentype"
)

type Drawable interface {
	Draw(tx *PrintTx) int
}

type PrintRequest struct {
	Drawings []Drawable
	Images   []ImageData
	Margin   float64
}

type PrintTx struct {
	MonoChrome      bool
	MonoChromeColor color.Color
	Rgba            draw.Image
	Src             image.Image
	Fg              image.Image
	Bg              *image.Uniform
	Dpi             float64
	Margin          float64
	Fonts           map[string]opentype.Font
	Faces           *map[string]font.Face
	Images          map[string]image.Image
}
