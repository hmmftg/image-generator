package imagegenerator

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"image"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
)

type ImageData struct {
	Dpi            float64
	Height         int
	Width          int
	BackgroundFile string
	NeedB64        bool
	FileFormat     string
	File           string
}

func getGzip(src []byte) ([]byte, error) {
	var gzipBuffer bytes.Buffer
	gz := gzip.NewWriter(&gzipBuffer)
	if _, err := gz.Write(src); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return gzipBuffer.Bytes(), nil
}

func getBmpB64(src image.Image) (string, error) {
	var imageBuffer bytes.Buffer
	err := bmp.Encode(&imageBuffer, src)
	if err != nil {
		return "", err
	}
	compressed, err := getGzip(imageBuffer.Bytes())
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(compressed), nil
}

func getPngB64(src image.Image) (string, error) {
	var imageBuffer bytes.Buffer
	err := png.Encode(&imageBuffer, src)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageBuffer.Bytes()), nil
}

func writeBmpInFile(src image.Image, fileName string) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = bmp.Encode(b, src)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

func writePngInFile(src image.Image, fileName string) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, src)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

const (
	Bmp = "bmp"
	Png = "png"
)

func GetImage(img *ImageData, resp map[string]string, rgba *image.RGBA) error {
	var err error
	switch img.FileFormat {
	case Bmp:
		if img.NeedB64 {
			resp[Bmp], err = getBmpB64(rgba)
			if err != nil {
				return err
			}
		}

		if len(img.File) > 0 {
			err = writeBmpInFile(rgba, img.File)
			if err != nil {
				return err
			}
		}
	case Png:
		if img.NeedB64 {
			resp[Png], err = getPngB64(rgba)
			if err != nil {
				return err
			}
		}

		if len(img.File) > 0 {
			err = writePngInFile(rgba, img.File)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (img ImageData) InitializeImage(bgMap map[string]image.Image, tx *PrintTx,
) {
	tx.Rgba = image.NewRGBA(image.Rect(0, 0, img.Width, img.Height))
	if tx.Src != nil {
		draw.Draw(tx.Rgba, tx.Rgba.Bounds(), tx.Src, image.Point{}, draw.Over)
	} else {
		draw.Draw(tx.Rgba, tx.Rgba.Bounds(), tx.Bg, image.Point{}, draw.Src)
	}
}
