package stencil

import (
	"image"
)

type Anchor string

const (
	TopCenter Anchor = "top"
	TopLeft   Anchor = "top-left"
	TopRight  Anchor = "top-right"

	BottomCenter Anchor = "bottom"
	BottomLeft   Anchor = "bottom-left"
	BottomRight  Anchor = "bottom-right"

	Center      Anchor = "center"
	LeftCenter  Anchor = "left"
	RightCenter Anchor = "right"
)

func (a Anchor) Valid() bool {
	switch a {
	case TopCenter, TopLeft, TopRight,
		BottomCenter, BottomLeft, BottomRight,
		Center, LeftCenter, RightCenter:
		return true
	default:
		return false
	}
}

func (a Anchor) Position(x, y, sizeX, sizeY int) image.Point {
	switch a {
	case TopCenter:
		return image.Point{
			X: -(x - sizeX/2),
			Y: -y,
		}
	default:
		return image.Point{
			X: -x,
			Y: -y,
		}
	case TopRight:
		return image.Point{
			X: -(x - sizeX),
			Y: -y,
		}

	case BottomCenter:
		return image.Point{
			X: -(x - sizeX/2),
			Y: -(y - sizeY),
		}
	case BottomLeft:
		return image.Point{
			X: -(x - sizeX),
			Y: -(y - sizeY),
		}
	case BottomRight:
		return image.Point{
			X: -x,
			Y: -(y - sizeY),
		}

	case Center:
		return image.Point{
			X: -(x - sizeX/2),
			Y: -(y - sizeY/2),
		}
	case LeftCenter:
		return image.Point{
			X: -x,
			Y: -(y - sizeY/2),
		}
	case RightCenter:
		return image.Point{
			X: -(x - sizeX),
			Y: -(y - sizeY/2),
		}
	}
}
