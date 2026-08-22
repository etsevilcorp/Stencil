package stencil

type Stencil struct {
	X      int    `toml:"x"`
	Y      int    `toml:"y"`
	Anchor Anchor `toml:"anchor"` // according to which point on image the X and Y will be applied

	MaxHeight   *int        `toml:"max-height"`
	MaxWidth    *int        `toml:"max-width"`
	MaxHandling MaxHandling `toml:"max-handling"` // what will happen to image if it's too big

	MinHeight   *int        `toml:"min-height"`
	MinWidth    *int        `toml:"min-width"`
	MinHandling MinHandling `toml:"min-handling"` // what will happen to image if it's too small

	ZIndex int `toml:"z-index"` // bigger z-index, the "higher" the image is applied

	// misc
	// rotation
	Perspective int   `toml:"perspective"`
	RotateX     int16 `toml:"rotate-x"` // deg clock-wise
	RotateY     int16 `toml:"rotate-y"` // a reminder that 2D images have no depth
	RotateZ     int16 `toml:"rotate-z"` // another reminder that 2D images have no depth(well, 0, actually)
}

type Stencils map[string]Stencil
