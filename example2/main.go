package main

import (
	"flag"
	"fmt"
	"log"

	ig "github.com/hmmftg/image-generator"
)

const borderSpace = 0.05

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
	imageX = flag.Int("imageX", 1000, "imageX")
	imageY = flag.Int("imageY", 1000, "imageY")
	text1X = flag.Float64("text1X", borderSpace*2, "textX")
	text1Y = flag.Float64("text1Y", borderSpace*5, "texty")
	text2X = flag.Float64("text2X", borderSpace*2, "textX")
	text2Y = flag.Float64("text2Y", borderSpace*15, "texty")

	text1 = string("سلام نوشته 11111 ۱۱۱۱")
	text2 = string("سلام نوشته 11111 ۱۱۱۱ sd;alsd asd سشمیبسمیبنتمسشیبتکمشسینتبکمشت ینمشسیتب شسیتب شسنمیاب شسیاکبا شکسیاب کشسیابشکسایبا")
	text3 = string("145282220.25")
	text4 = string("145282220.25")
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
		Margin: borderSpace * 2,
		Drawings: []ig.Drawable{
			ig.Text{
				Text:       text1,
				X:          *text1X,
				Y:          *text1Y,
				RightAlign: false,
				FontFace:   faceName,
			},
			ig.Text{
				Text:                 text3,
				X:                    *text1X + 0.1,
				Y:                    *text1Y + 0.1,
				RightAlign:           false,
				FontFace:             faceName,
				NumberToPersianWords: true,
			},
			ig.Text{
				Text:          text3,
				X:             *text1X + 0.2,
				Y:             *text1Y + 0.2,
				RightAlign:    false,
				FontFace:      faceName,
				NumberToWords: true,
			},
			ig.Text{
				Text:       text2,
				X:          *text2X,
				Y:          *text2Y,
				RightAlign: true,
				MaxWidth:   0.9,
				FontFace:   fmt.Sprintf("%s:%f", *fontfile, 17.85),
			},
			ig.Text{
				Text:            text2,
				X:               *text2X + borderSpace,
				Y:               *text2Y + borderSpace,
				RightAlign:      true,
				NumbersToArabic: true,
				MaxWidth:        0.9,
				FontFace:        faceName,
			},
			ig.Text{
				Text:       "l;asdj aklsdj ajsd lasjdl asjlkd asjd laksdljasl d lasjdl aslkdj lasjdl alsda",
				X:          *text2X,
				Y:          *text2Y + borderSpace*2,
				RightAlign: false,
				MaxWidth:   0.6,
				FontFace:   faceName,
			},
			ig.Text{
				Text:       "l;asdj aklsdj ajsd lasjdl asjlkd asjd laksdljasl d lasjdl aslkdj lasjdl alsda",
				X:          *text2X,
				Y:          *text2Y + borderSpace*3,
				RightAlign: false,
				MaxWidth:   0.7,
				FontFace:   faceName,
			},
			ig.Text{
				Text:       "l;asdj aklsdj ajsd lasjdl asjlkd asjd laksdljasl d lasjdl aslkdj lasjdl alsda",
				X:          *text2X,
				Y:          *text2Y + borderSpace*4,
				RightAlign: false,
				MaxWidth:   0.8,
				FontFace:   faceName,
			},
			ig.Rect{
				Thickness: 2,
				X1:        borderSpace,
				X2:        1 - borderSpace,
				Y1:        borderSpace,
				Y2:        1 - borderSpace,
				Color:     ig.GreenRuler,
			},
			ig.Rect{
				Thickness: 2,
				X1:        borderSpace * 2,
				X2:        1 - (borderSpace * 2),
				Y1:        borderSpace * 2,
				Y2:        1 - (borderSpace * 2),
				Color:     ig.RedRuler,
			},
			&ig.Image{
				ID:         "logo1.png",
				X:          borderSpace * 2,
				Y:          borderSpace * 2,
				RightAlign: true,
				Scale:      0.3,
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
