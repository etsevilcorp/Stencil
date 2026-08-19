package stencil

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
