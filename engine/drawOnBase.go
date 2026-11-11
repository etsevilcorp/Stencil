package engine

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/etsevilcorp/stencil/stencil"
)

type Layer struct {
	Name  string
	Cfg   stencil.Stencil
	Image image.Image
}

func Apply(base draw.Image, img image.Image, cfg stencil.Stencil) {
	img = cfg.Fit(img)
	b := img.Bounds()
	pt := cfg.Anchor.Position(cfg.X, cfg.Y, b.Dx(), b.Dy())
	draw.Draw(base, base.Bounds(), img, pt, draw.Over)
}

func Draw(base draw.Image, layers []Layer) {
	for _, l := range layers {
		Apply(base, l.Image, l.Cfg)
	}
}

func Solid(c color.Color, w, h int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), image.NewUniform(c), image.Point{}, draw.Src)
	return img
}
