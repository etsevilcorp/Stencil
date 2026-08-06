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

type MaxHandling string

const (
	Clip           MaxHandling = "clip"            // clips
	Compress       MaxHandling = "compress"        // compresses both sides until they fit
	CompressAspect MaxHandling = "compress-aspect" // compress while saving aspect ratio
)

type MinHandling string

const (
	Stretch       MinHandling = "stretch"        // stretches both sides until they fit
	StretchAspect MinHandling = "stretch-aspect" // stretch while saving aspect ratio
	Repeat        MinHandling = "repeat"         // repeats the same image, so the whole min will be filled
)

type Stencil struct {
	X      int    `toml:"x"`
	Y      int    `toml:"y"`
	Anchor Anchor `toml:"anchor"`

	MaxHeight   *int        `toml:"max-height"`
	MaxWidth    *int        `toml:"max-width"`
	MaxHandling MaxHandling `toml:"max-handling"`

	MinHeight   *int        `toml:"min-height"`
	MinWidth    *int        `toml:"min-width"`
	MinHandling MinHandling `toml:"min-handling"`

	// misc
	// rotation
	Perspective int   `toml:"perspective"`
	RotateX     int16 `toml:"rotate-x"` // deg clock-wise
	RotateY     int16 `toml:"rotate-y"` // a reminder that 2D images have no depth
	RotateZ     int16 `toml:"rotate-z"` // another reminder that 2D images have no depth(well, 0, actually)
}

type Stencils map[string]Stencil
