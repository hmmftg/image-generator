package imagegenerator

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"github.com/hmmftg/image/bmp"
	"github.com/hmmftg/image/draw"
)

type ImageData struct {
	Dpi             float64
	Height          int
	Width           int
	BackgroundFile  string
	NeedB64         bool
	FileFormat      string
	File            string
	MonoChrome      bool
	MonoChromeColor color.Color
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

func getMonoChromeBmpB64(src image.Image) (string, error) {
	var imageBuffer bytes.Buffer
	err := bmp.Encode(&imageBuffer, src)
	if err != nil {
		return "", err
	}
	var monoChromeBuffer bytes.Buffer
	err = getBmpMonoChrome(&imageBuffer, &monoChromeBuffer)
	if err != nil {
		return "", err
	}

	compressed, err := getGzip(monoChromeBuffer.Bytes())
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(compressed), nil
}

func getBmpBuffer(src image.Image) (io.Reader, error) {
	var imageBuffer bytes.Buffer
	err := bmp.Encode(&imageBuffer, src)
	if err != nil {
		return nil, err
	}
	return &imageBuffer, nil
}

func getBmpMonoChrome(buffer io.Reader, target io.Writer) error {
	err := EncodeMonoChrome(target, buffer)
	if err != nil {
		return err
	}
	return nil
}

func getPngB64(src image.Image) (string, error) {
	var imageBuffer bytes.Buffer
	err := png.Encode(&imageBuffer, src)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageBuffer.Bytes()), nil
}

func writeMonoChromeBMPInFile(src image.Image, fileName string) error {
	buffer, err := getBmpBuffer(src)
	if err != nil {
		return err
	}
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = getBmpMonoChrome(buffer, b)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
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
	Bmp           = "bmp"
	BmpMonoChrome = "bmp-monochrome"
	Png           = "png"
)

func GetImage(img *ImageData, resp map[string]string, rgba draw.Image) error {
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
	case BmpMonoChrome:
		if img.Dpi != 150 {
			return nil
		}
		if img.NeedB64 {
			resp[BmpMonoChrome], err = getMonoChromeBmpB64(rgba)
			if err != nil {
				return err
			}
		}

		if len(img.File) > 0 {
			err = writeMonoChromeBMPInFile(rgba, img.File)
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

func (tx *PrintTx) getNewImage(rect image.Rectangle) draw.Image {
	if tx.MonoChrome {
		return image.NewPaletted(rect, color.Palette{tx.MonoChromeColor})
	} else {
		return image.NewRGBA(rect)
	}
}

func (img ImageData) InitializeImage(bgMap map[string]image.Image, tx *PrintTx,
) {
	rect := image.Rect(0, 0, img.Width*(int(img.Dpi/72.)), img.Height*(int(img.Dpi/72.)))

	tx.Rgba = tx.getNewImage(rect)
	if tx.Src != nil {
		if tx.Src.Bounds() != rect {
			scaledImage := tx.getNewImage(rect)
			draw.NearestNeighbor.Scale(scaledImage, rect, tx.Src, tx.Src.Bounds(), draw.Over, nil)
			draw.Draw(tx.Rgba, tx.Rgba.Bounds(), scaledImage, image.Point{}, draw.Over)
		} else {
			draw.Draw(tx.Rgba, tx.Rgba.Bounds(), tx.Src, image.Point{}, draw.Over)
		}
	} else {
		draw.Draw(tx.Rgba, tx.Rgba.Bounds(), tx.Bg, image.Point{}, draw.Src)
	}
}
