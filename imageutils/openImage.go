package imageutils

import (
	"image"
	"os"
)

func OpenImage(path string) (image.Image, error) {
	stFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer stFile.Close()

	stenciling, _, err := image.Decode(stFile)
	if err != nil {
		return nil, err
	}

	return stenciling, nil
}
