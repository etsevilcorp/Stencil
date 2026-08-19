package stencil

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
