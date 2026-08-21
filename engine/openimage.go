package engine

import (
	"image"
	"os"
)

func openImage(path string) (image.Image, error) {
	stFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stenciling, _, err := image.Decode(stFile)
	if err != nil {
		return nil, err
	}

	return stenciling, nil
}
