package stencil

import (
	"image"
	"image/draw"
	"math"

	drawx "golang.org/x/image/draw"
)

type MaxHandling string

const (
	Clip           MaxHandling = "clip"            // clips
	Compress       MaxHandling = "compress"        // compresses both sides until they fit
	CompressAspect MaxHandling = "compress-aspect" // compress while saving aspect ratio
)

func (mh MaxHandling) Valid() bool {
	switch mh {
	case Clip, Compress, CompressAspect:
		return true
	default:
		return false
	}
}

// whenthe doesn't escape to heap

func (s Stencil) HandleMax(i image.Image) image.Image {
	switch s.MaxHandling {
	default:
		return s.clip(i)
	case Compress:
		return s.compress(i)
	case CompressAspect:
		return s.compressAspect(i)
	}
}

func (s Stencil) InMax(bounds image.Rectangle) bool {
	return (s.MaxHeight == nil || *s.MaxHeight >= bounds.Max.Y) && (s.MaxWidth == nil || *s.MaxWidth >= bounds.Max.X)
}

func (s Stencil) clip(i image.Image) image.Image {
	bounds := i.Bounds()
	var endX, endY = bounds.Dx(), bounds.Dy()
	if s.MaxHeight != nil && *s.MaxHeight < endY {
		endY = *s.MaxHeight
	}
	if s.MaxWidth != nil && *s.MaxWidth < endX {
		endX = *s.MaxWidth
	}

	interfaceSub, ok := i.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	if !ok {
		panic("Image doesn't support cropping(this wasn't in instruction)")
	}

	crRect := image.Rect(0, 0, endX, endY)
	cropped := interfaceSub.SubImage(crRect)
	newImage := image.NewRGBA(image.Rect(0, 0, crRect.Dx(), crRect.Dy()))
	draw.Draw(newImage, newImage.Bounds(), cropped, crRect.Min, draw.Over)
	return newImage
}

func (s Stencil) compress(i image.Image) image.Image {
	bounds := i.Bounds()
	var endX, endY = bounds.Dx(), bounds.Dy()
	if s.MaxHeight != nil && *s.MaxHeight < i.Bounds().Dy() {
		endY = *s.MaxHeight
	}
	if s.MaxWidth != nil && *s.MaxWidth < i.Bounds().Dx() {
		endX = *s.MaxWidth
	}

	dstRect := image.Rect(0, 0, endX, endY)
	newImg := image.NewRGBA(dstRect)
	drawx.CatmullRom.Scale(newImg, dstRect, i, bounds, draw.Over, nil)

	return newImg
}

func (s Stencil) compressAspect(i image.Image) image.Image {
	bounds := i.Bounds()
	var endX, endY = bounds.Dx(), bounds.Dy()

	var change float64 = 1
	if s.MaxHeight != nil && *s.MaxHeight < i.Bounds().Dy() {
		change = float64(*s.MaxHeight) / float64(endY)
	}
	if s.MaxWidth != nil && *s.MaxWidth < i.Bounds().Dx() {
		changeX := float64(*s.MaxWidth) / float64(endX)
		if changeX < change {
			change = changeX
		}
	}

	dstRect := image.Rect(0, 0, int(math.Round(float64(endX)*change)), int(math.Round(float64(endY)*change)))
	newImg := image.NewRGBA(dstRect)
	drawx.CatmullRom.Scale(newImg, dstRect, i, bounds, draw.Over, nil)

	return newImg
}
