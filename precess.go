package imagegenerator

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	BmpAndPng = "bmp+png"
	JustPng   = "png"
)

func ProcessRequest(
	requests map[string]PrintRequest,
	fonts map[string]opentype.Font,
	images map[string]image.Image,
) (map[string]string, error) {
	fullResp := make(map[string]string, 0)
	faces := make(map[string]font.Face, 0)

	for name, request := range requests {
		resp := make(map[string]string, 0)
		for imageID := range request.Images {
			tx := PrintTx{Faces: &faces, Dpi: request.Images[imageID].Dpi, Images: images, Fonts: fonts}
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
