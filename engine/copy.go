package engine

import "image"

// got proved wrong
// everything is possible with reflect

func CopyRGBA(src *image.RGBA) *image.RGBA {
	clone := *src
	clone.Pix = append([]byte(nil), src.Pix...)
	return &clone
}
