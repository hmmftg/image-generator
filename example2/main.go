package main

import (
	"flag"
	"fmt"
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
	fonts := map[string]*opentype.Font{
		*fontfile: f,
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

	var faceName = fmt.Sprintf("%s:%f", *fontfile, *size)

	req := ig.PrintRequest{
		Drawings: []ig.Drawable{
			ig.Text{
				Text:       text1,
				X:          *text1X,
				Y:          *text1Y,
				RightAlign: false,
				FontFace:   faceName,
			},
			ig.Text{
				Text:       text2,
				X:          *text2X,
				Y:          *text2Y,
				RightAlign: true,
				FontFace:   faceName,
			},
		},
		Images: []ig.ImageData{
			{
				Dpi:        *dpi,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Bmp,
				File:       "s1.bmp",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Png,
				File:       "s1.png",
				NeedB64:    false,
			},
		},
	}

	result, err := ig.ProcessRequest(map[string]ig.PrintRequest{"t1": req}, fonts, nil)
	fmt.Println(err, result)

}
