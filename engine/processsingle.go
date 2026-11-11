package engine

import (
	"fmt"
	"image/draw"
	"path/filepath"
	"sort"

	_ "image/jpeg"

	"github.com/etsevilcorp/stencil/imageutils"
	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils, out string) error {
	stencilsSorted := sortNames(stencils, cfgs)

	var nameRes = make(sort.StringSlice, 0, len(stencils))
	for _, name := range stencilsSorted {
		nameRes = append(nameRes, name)
	}
	nameRes.Sort()

	layers, err := LoadLayers(nameRes, cfgs, stencils)
	if err != nil {
		return fmt.Errorf("processsingle: unable to load layers: %w", err)
	}
	Draw(base, layers)

	for _, l := range layers {
		err := imageutils.CreateImage(filepath.Join(out, l.Name+".png"), base)
		if err != nil {
			return fmt.Errorf("processsingle couldn't create image %v", l.Name)
		}
	}

	return nil
}
