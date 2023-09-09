package imagegenerator

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
	"golang.org/x/image/font/opentype"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func LoadFonts(src string) (map[string]*opentype.Font, error) {
	if !exists(src) {
		return nil, nil
	}
	fontList, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	fonts := make(map[string]*opentype.Font, 0)
	for _, fontName := range fontList {
		fontBytes, err := os.ReadFile(filepath.Join(src, fontName.Name()))
		if err != nil {
			return nil, err
		}
		f, err := opentype.Parse(fontBytes)
		if err != nil {
			return nil, err
		}
		fonts[fontName.Name()] = f
	}
	return fonts, nil
}

func LoadImages(src string) (map[string]image.Image, error) {
	if !exists(src) {
		return nil, nil
	}
	imageList, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	fonts := make(map[string]image.Image, 0)
	for _, imageName := range imageList {
		imageFile, err := os.Open(filepath.Join(src, imageName.Name()))
		if err != nil {
			return nil, err
		}
		extention := filepath.Ext(imageFile.Name())
		var imageData image.Image
		switch extention {
		case ".png":
			imageData, err = png.Decode(imageFile)
		case ".bmp":
			imageData, err = bmp.Decode(imageFile)
		case ".jpg":
			imageData, err = jpeg.Decode(imageFile)
		}
		if err != nil {
			return nil, err
		}
		fonts[imageName.Name()] = imageData
	}
	return fonts, nil
}
