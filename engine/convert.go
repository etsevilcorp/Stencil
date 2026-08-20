package engine

import (
	"image"
	"image/draw"
)

func ConvertToDrawImage(src image.Image) draw.Image {
	if dimg, ok := src.(draw.Image); ok {
		return dimg
	}

	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)

	return dst
}
