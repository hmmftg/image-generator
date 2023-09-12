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
	X                float64
	Y                float64
	MaxWidth         float64 // if result width is more than this value, library tries new font face with decreased size
	RightAlign       bool
	NumbersToArabic  bool
	NumbersToPersian bool
	FontFace         string
}

func (tx *PrintTx) AddFace(fontName string, sz float64) font.Face {
	name := fmt.Sprintf("%s:%f:%f", fontName, sz, tx.Dpi)
	face, ok := (*tx.Faces)[name]
	if ok {
		return face
	}
	opts := opentype.FaceOptions{
		Size:    sz,
		DPI:     tx.Dpi,
		Hinting: font.HintingFull,
	}
	opts.Hinting = font.HintingFull

	fn := tx.Fonts[fontName]
	face, _ = opentype.NewFace(&fn, &opts)
	// fmt.Println("Adding font face:", name)
	(*tx.Faces)[name] = face
	return face
}

func (s Text) CheckFace(tx *PrintTx) font.Face {
	face, ok := (*tx.Faces)[s.Font(tx.Dpi)]
	if !ok {
		faceName := s.Font(tx.Dpi)
		if len(faceName) == 0 {
			log.Println("invalid font face(text ignored)", s.FontFace, tx.Dpi)
			return nil
		}
		return tx.AddFace(s.FontData())
	}
	return face
}

func (s Text) Font(dpi float64) string {
	return fmt.Sprintf("%s:%f", s.FontFace, dpi)
}

func (s Text) FontData() (string, float64) {
	fn := strings.Split(s.FontFace, ":")
	if len(fn) == 0 {
		return "", 0
	}
	sz, err := strconv.ParseFloat(fn[1], 64)
	if err != nil {
		return "", 0
	}
	return fn[0], sz
}

func (s Text) Draw(tx *PrintTx) int {
	// log.Println("drawing", s)
	firstRune := []rune(s.Text)[0]
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
	x := tx.RelationalX(s.X)
	y := tx.RelationalY(s.Y)

	d := &font.Drawer{
		Dst:  tx.Rgba,
		Src:  tx.Bg,
		Face: face,
		Dot:  fixed.P(x, y),
	}
	textLen := d.MeasureString(s.Text)
	advance := textLen.Round() + x
	if s.MaxWidth > 0 {
		for i := 0.0; tx.CoordinationX(advance) > s.MaxWidth; i += 0.05 {
			fn, sz := s.FontData()
			d.Face = tx.AddFace(fn, sz-i)
			textLen = d.MeasureString(s.Text)
			advance = textLen.Round() + x
		}
	}
	for advance > tx.Rgba.Bounds().Max.X-tx.RelationalX(tx.Margin) {
		if garabic.IsArabicLetter(firstRune) {
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
	// fmt.Println("drawing", advance, tx.Dpi, d.Dot, s.Text)

	d.DrawString(s.Text)
	return advance
}
