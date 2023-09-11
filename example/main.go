package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	ig "github.com/hmmftg/image-generator"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const borderSpace = 0.05

var (
	dpi        = flag.Float64("dpi", 300, "screen resolution in Dots Per Inch")
	basePath   = flag.String("base", "..", "filename of the ttf font")
	fontFolder = flag.String("fonts", "fonts", "name of the ttfs folder")
	fontfile   = flag.String("fontfile", "arial.ttf", "filename of the ttf font")
	hinting    = flag.String("hinting", "full", "none | full")
	size       = flag.Float64("size", 20, "font size in points")
	imageX     = flag.Int("imageX", 800, "imageX")
	imageY     = flag.Int("imageY", 500, "imageY")
	text1X     = flag.Float64("text1X", borderSpace*2, "textX")
	text1Y     = flag.Float64("text1Y", borderSpace*10, "texty")
	text2X     = flag.Float64("text2X", borderSpace*2, "textX")
	text2Y     = flag.Float64("text2Y", borderSpace*30, "texty")

	text1 = string("سلام نوشته")
	text2 = string("سلام نوشته")
)

func drawBorder(tx *ig.PrintTx, thickness int) {
	ig.Rect{Thickness: thickness, X1: borderSpace, Y1: borderSpace, X2: -1 * borderSpace, Y2: -1 * borderSpace, Color: ig.BlackRuler}.Draw(tx)
}

func main() {
	flag.Parse()
	fmt.Printf("Loading fontfile %q %q %q\n", *basePath, *fontFolder, *fontfile)
	b, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", *basePath, *fontFolder, *fontfile))
	if err != nil {
		log.Println(err)
		return
	}
	f, err := opentype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	// Truetype stuff
	opts := opentype.FaceOptions{
		Size:    *size,
		DPI:     *dpi,
		Hinting: font.HintingFull,
	}
	switch *hinting {
	default:
		opts.Hinting = font.HintingNone
	case "full":
		opts.Hinting = font.HintingFull
	}
	fonts := map[string]opentype.Font{*fontfile: *f}
	// Truetype stuff
	face, err := opentype.NewFace(f, &opts)
	if err != nil {
		log.Fatal(err)
	}
	var faceName = fmt.Sprintf("%s:%f", *fontfile, *size)
	faces := map[string]font.Face{faceName: face}
	// Freetype context
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, *imageX, *imageY))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	tx := ig.PrintTx{Rgba: rgba, Bg: fg, Faces: &faces, Dpi: *dpi, Fonts: fonts}

	// Make some background

	// Draw the guidelines.
	drawBorder(&tx, 3)

	ig.Line{X1: borderSpace * 2, X2: borderSpace * 2, Y1: borderSpace * 2, Y2: borderSpace * 2.5, Thickness: 1, Color: ig.GreenRuler}.Draw(&tx)
	offset1 := ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X + 50, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X + 100, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y + 50, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y + 100, RightAlign: false}.Draw(&tx)
	ig.Line{X1: float64(offset1) / float64(*imageX), X2: borderSpace * 2, Y1: borderSpace * 2.5, Y2: borderSpace * 2.5, Thickness: 1, Color: ig.GreenRuler}.Draw(&tx)
	ig.Line{X1: -1 * borderSpace * 2, X2: borderSpace * 2.5, Y1: borderSpace * 2.5, Y2: -1 * borderSpace * 2, Thickness: 1, Color: ig.GreenRuler}.Draw(&tx)
	offset2 := ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X + 50, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X + 50 + 50, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y + 50, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y + 50 + 50, RightAlign: true}.Draw(&tx)
	ig.Line{X1: float64(offset2) / float64(*imageX), X2: -1 * borderSpace * 2, Y1: -1 * borderSpace * 2, Y2: -1 * borderSpace * 2, Thickness: 1, Color: ig.GreenRuler}.Draw(&tx)

	//HLine(borderSpace*2, borderSpace*27.5, offset.X.Ceil(), rgba, GreenRuler)

	// Save that RGBA image to disk.
	outFile, err := os.Create("s4.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	bf := bufio.NewWriter(outFile)
	err = png.Encode(bf, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = bf.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")

}
