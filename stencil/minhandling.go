package stencil

import (
	"image"
	"image/draw"
	"math"

	drawx "golang.org/x/image/draw"
)

type MinHandling string

const (
	Stretch       MinHandling = "stretch"        // stretches both sides until they fit
	StretchAspect MinHandling = "stretch-aspect" // stretch while saving aspect ratio
	Repeat        MinHandling = "repeat"         // repeats the same image, so the whole min will be filled
)

func (mh MinHandling) Valid() bool {
	switch mh {
	case Stretch, StretchAspect, Repeat:
		return true
	default:
		return false
	}
}

func (s Stencil) HandleMin(i image.Image) image.Image {
	switch s.MinHandling {
	default:
		bounds := i.Bounds()
		var endX, endY = bounds.Dx(), bounds.Dy()
		if s.MinHeight != nil && *s.MinHeight > endY {
			endY = *s.MinHeight
		}
		if s.MinWidth != nil && *s.MinWidth > endX {
			endX = *s.MinWidth
		}

		crRect := image.Rect(0, 0, endX, endY)
		newImage := image.NewRGBA(image.Rect(0, 0, crRect.Dx(), crRect.Dy()))

		drawx.CatmullRom.Scale(newImage, crRect, i, bounds, draw.Src, nil)
		return newImage
	case Repeat:
		bounds := i.Bounds()
		var endX, endY = bounds.Dx(), bounds.Dy()
		if s.MinHeight != nil && *s.MinHeight > endY {
			endY = *s.MinHeight
		}
		if s.MinWidth != nil && *s.MinWidth > endX {
			endX = *s.MinWidth
		}

		dstRect := image.Rect(0, 0, endX, endY)
		newImg := image.NewRGBA(dstRect)

		var sizeX, sizeY = bounds.Dx(), bounds.Dy()

		// i already taken
		for row := range int(math.Ceil(float64(endX) / float64(bounds.Dx()))) {
			for col := range int(math.Ceil(float64(endY) / float64(bounds.Dy()))) {
				draw.Draw(newImg, dstRect, i, image.Pt(-(sizeX*row), -(sizeY*col)), draw.Src)
			}
		}

		return newImg
	case StretchAspect:
		bounds := i.Bounds()
		var endX, endY = bounds.Dx(), bounds.Dy()

		var change float64 = 1
		if s.MinHeight != nil && *s.MinHeight < i.Bounds().Dy() {
			change = float64(*s.MinHeight) / float64(endY)
		}
		if s.MinWidth != nil && *s.MinWidth < i.Bounds().Dx() {
			if changeX := float64(*s.MinWidth) / float64(endX); changeX > change {
				change = changeX
			}
		}

		dstRect := image.Rect(0, 0, int(math.Round(float64(endX)*change)), int(math.Round(float64(endY)*change)))
		newImg := image.NewRGBA(dstRect)
		drawx.CatmullRom.Scale(newImg, dstRect, i, bounds, draw.Src, nil)

		return newImg
	}
}
