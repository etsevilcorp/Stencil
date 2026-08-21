package engine

import (
	"image/draw"
	"image/png"
	"os"
)

func createImage(path string, base draw.Image) error {
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
