package imageutils

import (
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

func CreateImage(path string, base draw.Image) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, base)
	if err != nil {
		return err
	}

	return nil
}
