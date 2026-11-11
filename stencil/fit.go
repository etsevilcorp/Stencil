package stencil

import "image"

func (s Stencil) Fit(i image.Image) image.Image {
	if !s.InMax(i.Bounds()) {
		i = s.HandleMax(i)
	}
	if s.OutMin(i.Bounds()) {
		i = s.HandleMin(i)
	}
	return i
}
