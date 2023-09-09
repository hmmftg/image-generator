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

const borderSpace = 10

var (
	dpi        = flag.Float64("dpi", 300, "screen resolution in Dots Per Inch")
	fontFolder = flag.String("fonts", "./example/", "filename of the ttf font")
	fontfile   = flag.String("fontfile", "arial.ttf", "filename of the ttf font")
	hinting    = flag.String("hinting", "full", "none | full")
	size       = flag.Float64("size", 20, "font size in points")
	imageX     = flag.Int("imageX", 800, "imageX")
	imageY     = flag.Int("imageY", 500, "imageY")
	text1X     = flag.Int("text1X", borderSpace*2, "textX")
	text1Y     = flag.Int("text1Y", 100, "texty")
	text2X     = flag.Int("text2X", borderSpace*2, "textX")
	text2Y     = flag.Int("text2Y", 300, "texty")

	text1 = string("سلام نوشته")
	text2 = string("سلام نوشته")
)

func drawBorder(rgba *image.RGBA, thickness int) {
	ig.Rect{BackGround: rgba, Thickness: thickness, X1: borderSpace, Y1: borderSpace, X2: *imageX - borderSpace, Y2: *imageY - borderSpace, Color: ig.BlackRuler}.Draw()
}

func main() {
	flag.Parse()
	fmt.Printf("Loading fontfile %q %q\n", *fontFolder, *fontfile)
	b, err := os.ReadFile(*fontFolder + *fontfile)
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
	// Truetype stuff
	face, err := opentype.NewFace(f, &opts)
	if err != nil {
		log.Fatal(err)
	}
	var faceName = fmt.Sprintf("%s:%f", *fontfile, *size)
	fonts := map[string]font.Face{faceName: face}
	// Freetype context
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, *imageX, *imageY))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	tx := ig.PrintTx{Rgba: rgba, Bg: fg, Fonts: &fonts}

	// Make some background

	// Draw the guidelines.
	drawBorder(rgba, 3)

	ig.VLine(rgba, ig.GreenRuler, borderSpace*2, borderSpace*2, *imageY-250)
	offset1 := ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X + 50, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X + 100, Y: *text1Y, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y + 50, RightAlign: false}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text1, X: *text1X, Y: *text1Y + 100, RightAlign: false}.Draw(&tx)
	ig.VLine(rgba, ig.GreenRuler, offset1, borderSpace*2, *imageY-250)
	ig.VLine(rgba, ig.RedRuler, *imageX-(borderSpace*2), *imageY-250, *imageY-(borderSpace*2))
	offset2 := ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X + 50, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X + 50 + 50, Y: *text2Y, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y + 50, RightAlign: true}.Draw(&tx)
	ig.Text{FontFace: faceName, Text: text2, X: *text2X, Y: *text2Y + 50 + 50, RightAlign: true}.Draw(&tx)
	ig.VLine(rgba, ig.RedRuler, offset2, *imageY-250, *imageY-(borderSpace*2))

	//HLine(borderSpace*2, borderSpace*27.5, offset.X.Ceil(), rgba, GreenRuler)

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
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
