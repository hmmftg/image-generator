package imagegenerator

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hmmftg/garabic"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type Text struct {
	Text             string
	X, Y             int
	RightAlign       bool
	NumbersToArabic  bool
	NumbersToPersian bool
	FontFace         string
}

func (s Text) CheckFace(tx *PrintTx) font.Face {
	face, ok := (*tx.Faces)[s.Font(tx.Dpi)]
	if !ok {
		faceName := s.Font(tx.Dpi)
		if len(faceName) == 0 {
			return nil
		}
		fmt.Println("Adding font face:", faceName)
		fn := strings.Split(faceName, ":")
		sz, _ := strconv.ParseFloat(fn[1], 64)
		opts := opentype.FaceOptions{
			Size:    sz,
			DPI:     tx.Dpi,
			Hinting: font.HintingFull,
		}
		opts.Hinting = font.HintingFull
		font := tx.Fonts[fn[0]]
		face, _ := opentype.NewFace(&font, &opts)
		(*tx.Faces)[faceName] = face
		return face
	}
	return face
}

func (s Text) Font(dpi float64) string {
	return fmt.Sprintf("%s:%f", s.FontFace, dpi)
}

func (s Text) Draw(tx *PrintTx) int {
	s.Text = garabic.Shape(s.Text)
	if s.NumbersToArabic {
		s.Text = garabic.ToArabicDigits(s.Text)
	}
	if s.NumbersToPersian {
		s.Text = garabic.ToPersianDigits(s.Text)
	}

	face := s.CheckFace(tx)
	if face == nil {
		log.Printf("face not detected")
		return 0
	}

	d := &font.Drawer{
		Dst:  tx.Rgba,
		Src:  tx.Bg,
		Face: face,
		Dot:  fixed.P(s.X, s.Y),
	}
	len := d.MeasureString(s.Text)
	advance := len.Round() + s.X
	if s.RightAlign {
		d.Dot = fixed.P(tx.Rgba.Bounds().Max.X-len.Round()-s.X, s.Y)
		advance = tx.Rgba.Bounds().Max.X - len.Round() - s.X
	}
	// fmt.Printf("Draw(%+v %T %+v)\n", s, *tx.Bg, d)

	d.DrawString(s.Text)

	// writePngInFile(tx.Rgba, fmt.Sprintf("%s_%d_%d.png", s.Text, s.X, s.Y))

	return advance
}
