package imagegenerator

import (
	"image"
	"image/color"
)

// HLine draws a horizontal line
func HLine(img *image.RGBA, col color.Color, x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, col color.Color, x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func DrawRect(x1, y1, x2, y2, thickness int, img *image.RGBA, col color.Color) {
	for i := 0; i < thickness; i++ {
		HLine(img, col, x1, y1+i, x2)
		HLine(img, col, x1, y2+i, x2+thickness-1)
		VLine(img, col, x1+i, y1, y2)
		VLine(img, col, x2+i, y1, y2)
	}
}

var (
	BlackRuler = color.RGBA{0x00, 0x00, 0x00, 0xff}
	GreenRuler = color.RGBA{0x40, 0xb0, 0xa0, 0xff}
	RedRuler   = color.RGBA{0xB0, 0x40, 0x40, 0xff}
)

type Rect struct {
	BackGround *image.RGBA
	Thickness  int
	X1, X2     int
	Y1, Y2     int
	Color      color.Color
}

func (r Rect) Draw() {
	DrawRect(r.X1, r.Y1, r.X2, r.Y2, r.Thickness, r.BackGround, r.Color)
}
