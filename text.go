package imagegenerator

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hmmftg/garabic"
	"github.com/hmmftg/image/font"
	"github.com/hmmftg/image/font/opentype"
	"github.com/hmmftg/image/math/fixed"
)

type Text struct {
	Text             string
	X, Y             float64
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
			log.Println("invalid font face(text ignored)", s.FontFace, tx.Dpi)
			return nil
		}
		// fmt.Println("Adding font face:", faceName)
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
	// log.Println("drawing", s)
	s.Text = garabic.Shape(s.Text)
	if s.NumbersToArabic {
		s.Text = garabic.ToArabicDigits(s.Text)
	}
	if s.NumbersToPersian {
		s.Text = garabic.ToPersianDigits(s.Text)
	}

	face := s.CheckFace(tx)
	if face == nil {
		return 0
	}
	x := tx.GetRelationalX(s.X)
	y := tx.GetRelationalY(s.Y)

	d := &font.Drawer{
		Dst:  tx.Rgba,
		Src:  tx.Bg,
		Face: face,
		Dot:  fixed.P(x, y),
	}
	textLen := d.MeasureString(s.Text)
	advance := textLen.Round() + x
	for advance > tx.Rgba.Bounds().Max.X-tx.GetRelationalX(tx.Margin) {
		if garabic.IsArabicLetter([]rune(s.Text)[0]) {
			s.Text = s.Text[1:]
		} else {
			s.Text = s.Text[:len(s.Text)-1]
		}
		textLen = d.MeasureString(s.Text)
		advance = textLen.Round() + x
	}
	if s.RightAlign {
		d.Dot = fixed.P(tx.Rgba.Bounds().Max.X-textLen.Round()-x, y)
		advance = tx.Rgba.Bounds().Max.X - textLen.Round() - x
	}
	// log.Println("drawing", d.Dot)

	d.DrawString(s.Text)
	return advance
}
