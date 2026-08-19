package stencil

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
