package imagegenerator

import (
	"image"
	"image/color"

	"github.com/StephaneBunel/bresenham"
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
	Thickness int
	X1, X2    float64
	Y1, Y2    float64
	Color     color.Color
}

func (r Rect) Draw(tx *PrintTx) int {
	DrawRect(tx.GetRelationalX(r.X1), tx.GetRelationalY(r.Y1), tx.GetRelationalX(r.X2), tx.GetRelationalY(r.Y2), r.Thickness*(int(tx.Dpi/72.)), tx.Rgba, r.Color)
	return tx.GetRelationalX(r.X2)
}

type Line struct {
	X1, X2    float64
	Y1, Y2    float64
	Thickness int
	Color     color.Color
}

func (l Line) Draw(tx *PrintTx) int {
	x1, y1, x2, y2 := tx.GetRelationalX(l.X1), tx.GetRelationalY(l.Y1), tx.GetRelationalX(l.X2), tx.GetRelationalY(l.Y2)
	for i := 0; i < l.Thickness; i++ {
		bresenham.DrawLine(tx.Rgba, x1, y1, x2, y2, l.Color)
		if x1 != x2 {
			x1 = x1 - 1
			x2 = x2 - 1
		}
		if y1 != y2 {
			y1 = y1 - 1
			y2 = y2 - 1
		}
	}
	return x2
}
