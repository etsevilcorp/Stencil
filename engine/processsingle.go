package engine

import (
	"fmt"
	"image/draw"
	"path/filepath"
	"sort"
	"strings"

	_ "image/jpeg"

	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils, out string) error {
	stencilsSorted := sortNames(stencils, cfgs)

	var nameRes = make(sort.StringSlice, 0, len(stencils))
	for _, name := range stencilsSorted {
		cfg := cfgs[name]
		err := cfg.Validate()
		if err != nil {
			return fmt.Errorf("processsingle failed to validate config: %w", err)
		}

		path := stencils[name]

		stPath := filepath.Join(name, path)
		stenciling, err := openImage(stPath)
		if err != nil {
			return fmt.Errorf("processsingle failed to process image %v: %w", stPath, err)
		}

		nameRes = append(nameRes, (strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))))

		bounds := stenciling.Bounds()
		if !cfg.InMax(bounds) {
			stenciling = cfg.HandleMax(stenciling)
		}
		if !cfg.OutMin(bounds) {
			stenciling = cfg.HandleMin(stenciling)
		}
		bounds = stenciling.Bounds() // after minmax

		draw.Draw(base, base.Bounds(), stenciling, cfg.Anchor.Position(cfg.X, cfg.Y, bounds.Dx(), bounds.Dy()), draw.Over)
	}
	nameRes.Sort()

	newName := fmt.Sprintf("%v-%v", baseName, strings.Join(nameRes, "-")) + ".png"
	err := createImage(filepath.Join(out, newName), base)
	if err != nil {
		return fmt.Errorf("processsingle couldn't create image ")
	}

	return nil
}
