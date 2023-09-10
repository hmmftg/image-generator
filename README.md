# image-generator
Library for generate image from text structure

all measures are pixel/inch

## example:
```go
package main

import (
	"flag"
	"fmt"
	"log"

	ig "github.com/hmmftg/image-generator"
)

var (
	basePath     = flag.String("base", "..", "base path")
	imagesFolder = flag.String("images", "images", "name of the images folder")
	fontFolder   = flag.String("fonts", "fonts", "name of the fonts folder")
	fontfile     = flag.String("fontfile", "arial.ttf", "filename of the ttf font")
	imageWidth   = flag.Int("width", 800, "image width")
	imageHeight  = flag.Int("height", 800, "image height")
)

func main() {
	flag.Parse()

    // first load ttf font files
	fonts, err := ig.LoadFonts(fmt.Sprintf("%s/%s", *basePath, *fontFolder))
	if err != nil {
		log.Fatal(err)
	}

    // load asset images(used as background or logo or etc)
	images, err := ig.LoadImages(fmt.Sprintf("%s/%s", *basePath, *imagesFolder))
	if err != nil {
		log.Fatal(err)
	}

	req := ig.PrintRequest{ // each request contains multiple objects to be printed on all defined outputs  
		Drawings: []ig.Drawable{ // each drawing is a printable element, currently supporting: text, rectangle and image
			ig.Text{ // print text on dimension
				Text:       "Hello World!",
				X:          4,
				Y:          4,
				RightAlign: false, // rtl
				FontFace:   "arial.ttf:30", // font name(should exist in defined fonts folder) : font size
			},
			ig.Rect{ // print rectangle on dimesion
				Thickness: 2,
				X1:        8,
				X2:        *imageWidth - 8,
				Y1:        8,
				Y2:        *imageHeight - 8,
				Color:     ig.GreenRuler,
			},
			&ig.Image{ // print image asset, like logo on dimesion
				ID: "logo1.png", // image asset which should exist in defined images folder
				X:  30,
				Y:  30,
			},
		},
		Images: []ig.ImageData{
			{
				Dpi:            300,
				BackgroundFile: "bg1.png", // background image, if nil white background used by default
				Height:         *imageY,
				Width:          *imageX,
				FileFormat:     ig.Png, // output file type(bmp and png supported)
				File:           "sB.png", // output file name
				NeedB64:        false, // will return base64 string if true
			},
			{
				Dpi:        72,
				Height:     *imageY,
				Width:      *imageX,
				FileFormat: ig.Png,
				File:       "s1.png",
				NeedB64:    false,
			},
		},
	}

    // for each request generate each image and save it and return result containing base64 of images if NeedB64 of them was true
	result, err := ig.ProcessRequest(map[string]ig.PrintRequest{"reuqest1": req}, fonts, images)
}
```