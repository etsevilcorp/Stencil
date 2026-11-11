package engine

import (
	"fmt"
	"path/filepath"

	"github.com/etsevilcorp/stencil/imageutils"
	"github.com/etsevilcorp/stencil/stencil"
)

func LoadLayers(order []string, cfgs stencil.Stencils, picks map[string]string) ([]Layer, error) {
	layers := make([]Layer, 0, len(order))
	for _, name := range order {
		cfg := cfgs[name]
		if err := cfg.Validate(); err != nil {
			return nil, fmt.Errorf("layer %q: %w", name, err)
		}
		p := filepath.Join(name, picks[name])
		img, err := imageutils.OpenImage(p)
		if err != nil {
			return nil, fmt.Errorf("layer %q: open %v: %w", name, p, err)
		}
		layers = append(layers, Layer{Name: name, Cfg: cfg, Image: img})
	}
	return layers, nil
}
