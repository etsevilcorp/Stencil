package engine

import (
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg"
	"image/png"

	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils) error {
	var nameRes strings.Builder
	nameRes.WriteString(baseName)
	nameRes.WriteRune('-')

	for name, path := range stencils {
		stFile, err := os.Open(filepath.Join(name, path))
		if err != nil {
			return err
		}
		_, err = nameRes.WriteString(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
		if err != nil {
			return err
		}
		nameRes.WriteRune('-')

		stenciling, _, err := image.Decode(stFile)
		if err != nil {
			return err
		}

		draw.Draw(base, base.Bounds(), stenciling, image.Pt(cfgs[name].X, cfgs[name].Y), draw.Src)
	}

	outputFile, err := os.Create(strings.TrimSuffix(nameRes.String(), "-") + ".png")
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
