package imagegenerator

import (
	"image"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	BmpAndPng = "bmp+png"
	JustPng   = "png"
)

func CheckFace(d Drawable, i ImageData, faces map[string]font.Face, fonts map[string]*opentype.Font) {
	if _, ok := faces[d.Font()]; !ok {
		fn := strings.Split(d.Font(), ":")
		sz, _ := strconv.Atoi(fn[1])
		opts := opentype.FaceOptions{
			Size:    float64(sz),
			DPI:     i.Dpi,
			Hinting: font.HintingFull,
		}
		opts.Hinting = font.HintingFull
		face, _ := opentype.NewFace(fonts[fn[0]], &opts)
		faces[d.Font()] = face
	}
}

func ProcessRequest(
	requests map[string]PrintRequest,
	fonts map[string]*opentype.Font,
	images map[string]image.Image,
) (map[string]string, error) {
	fullResp := make(map[string]string, 0)
	faces := make(map[string]font.Face, 0)

	for name, request := range requests {
		resp := make(map[string]string, 0)
		for imageID := range request.Images {
			tx := PrintTx{Fonts: &faces}
			if len(request.Images[imageID].BackgroundFile) > 0 {
				tx.Bg = image.Transparent
				tx.Src = images[request.Images[imageID].BackgroundFile]
			} else {
				tx.Bg = image.White
				tx.Src = nil
			}

			request.Images[imageID].InitializeImage(images, &tx)
			tx.Bg = image.Black

			// Draw the text.
			for id := range request.Drawings {
				CheckFace(request.Drawings[id], request.Images[imageID], faces, fonts)
				request.Drawings[id].Draw(&tx)
			}

			err := GetImage(&request.Images[imageID], resp, tx.Rgba)
			if err != nil {
				return resp, err
			}
		}

		switch name {
		case BmpAndPng:
			fullResp[Bmp] = resp[Bmp]
			fullResp[Png] = resp[Png]
		default:
			fullResp[name] = resp[Png]
		}
	}

	return fullResp, nil
}
