package main

import (
	"flag"
	"fmt"
	"log"

	ig "github.com/hmmftg/image-generator"
)

const borderSpace = 10

var (
	dpi1         = flag.Float64("dpi1", 72, "screen resolution in Dots Per Inch")
	dpi2         = flag.Float64("dpi2", 150, "screen resolution in Dots Per Inch")
	dpi3         = flag.Float64("dpi3", 300, "screen resolution in Dots Per Inch")
	basePath     = flag.String("base", "..", "base path")
	imagesFolder = flag.String("images", "images", "name of the images folder")
	fontFolder   = flag.String("fonts", "fonts", "name of the fonts folder")
	fontfile     = flag.String("fontfile", "arial.ttf", "filename of the ttf font")
	// hinting    = flag.String("hinting", "full", "none | full")
	size   = flag.Float64("size", 30, "font size in points")
	imageX = flag.Int("imageX", 800, "imageX")
	imageY = flag.Int("imageY", 800, "imageY")
	text1X = flag.Int("text1X", borderSpace*2, "textX")
	text1Y = flag.Int("text1Y", 100, "texty")
	text2X = flag.Int("text2X", borderSpace*2, "textX")
	text2Y = flag.Int("text2Y", 300, "texty")

	text1 = string("سلام نوشته 11111 ۱۱۱۱")
	text2 = string("سلام نوشته 11111 ۱۱۱۱")
)

func main() {
	flag.Parse()

	fonts, err := ig.LoadFonts(fmt.Sprintf("%s/%s", *basePath, *fontFolder))
	if err != nil {
		log.Fatal(err)
	}

	var faceName = fmt.Sprintf("%s:%f", *fontfile, *size)

	images, err := ig.LoadImages(fmt.Sprintf("%s/%s", *basePath, *imagesFolder))
	if err != nil {
		log.Fatal(err)
	}

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
			ig.Text{
				Text:            text2,
				X:               *text2X + borderSpace,
				Y:               *text2Y + borderSpace,
				RightAlign:      true,
				NumbersToArabic: true,
				FontFace:        faceName,
			},
			ig.Rect{
				Thickness: 2,
				X1:        borderSpace - 2,
				X2:        *imageX - borderSpace + 2,
				Y1:        borderSpace - 2,
				Y2:        *imageY - borderSpace + 2,
				Color:     ig.GreenRuler,
			},
			ig.Rect{
				Thickness: 2,
				X1:        borderSpace + 2,
				X2:        *imageX - borderSpace - 2,
				Y1:        borderSpace + 2,
				Y2:        *imageY - borderSpace - 2,
				Color:     ig.RedRuler,
			},
			&ig.Image{
				ID: "logo1.png",
				X:  borderSpace + 20,
				Y:  borderSpace + 20,
			},
		},
		Images: []ig.ImageData{
			{
				Dpi:            *dpi1,
				BackgroundFile: "bg1.png",
				Height:         *imageY,
				Width:          *imageX,
				FileFormat:     ig.Png,
				File:           "sB.png",
				NeedB64:        false,
			},
			{
				Dpi:        *dpi1,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Png,
				File:       "s1.png",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi2,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Png,
				File:       "s2.png",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi3,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Png,
				File:       "s3.png",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi1,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Bmp,
				File:       "s1.bmp",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi2,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Bmp,
				File:       "s2.bmp",
				NeedB64:    false,
			},
			{
				Dpi:        *dpi3,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Bmp,
				File:       "s3.bmp",
				NeedB64:    false,
			},
		},
	}

	result, err := ig.ProcessRequest(map[string]ig.PrintRequest{"t1": req}, fonts, images)
	fmt.Println(err, result)

}
