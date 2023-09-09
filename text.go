package imagegenerator

import (
	"fmt"

	"github.com/hmmftg/garabic"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Text struct {
	Text       string
	X, Y       int
	RightAlign bool
	FontFace   string
}

func (s Text) Font() string {
	return s.FontFace
}

func (s Text) Draw(tx *PrintTx) int {
	fmt.Printf("Draw(%s %+v %d %d)\n", s.Text, tx, s.X, s.Y)
	s.Text = garabic.Shape(s.Text)

	d := &font.Drawer{
		Dst:  tx.Rgba,
		Src:  tx.Bg,
		Face: (*tx.Fonts)[s.FontFace],
		Dot:  fixed.P(s.X, s.Y),
	}
	len := d.MeasureString(s.Text)
	advance := len.Round() + s.X
	if s.RightAlign {
		d.Dot = fixed.P(tx.Rgba.Bounds().Max.X-len.Round()-s.X, s.Y)
		advance = tx.Rgba.Bounds().Max.X - len.Round() - s.X
	}

	d.DrawString(s.Text)

	//writePngInFile(tx.Rgba, fmt.Sprintf("%s_%d_%d.png", s.Text, s.X, s.Y))

	return advance
}
