package engine

import (
	"image"
	"image/draw"
	"os"

	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, stencilFile *os.File, cfg stencil.Stencil) error {
	stenciling, _, err := image.Decode(stencilFile)
	if err != nil {
		return err
	}

	draw.Draw(base, base.Bounds(), stenciling, image.Pt(cfg.X, cfg.Y), draw.Src)
	return nil
}
