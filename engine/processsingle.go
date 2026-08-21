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

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils) error {
	stencilsSorted := sortNames(stencils, cfgs)

	var nameRes = make(sort.StringSlice, 0, len(stencils))
	for _, name := range stencilsSorted {
		cfg := cfgs[name]
		err := cfg.Validate()
		if err != nil {
			return err
		}

		path := stencils[name]

		stenciling, err := openImage(filepath.Join(name, path))
		if err != nil {
			return err
		}

		nameRes = append(nameRes, (strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))))

		bounds := stenciling.Bounds()
		if !cfg.InMax(bounds) {
			stenciling = cfg.HandleMax(stenciling)
		}
		if !cfg.OutMin(bounds) {
			stenciling = cfg.HandleMin(stenciling)
		}

		draw.Draw(base, base.Bounds(), stenciling, cfg.Anchor.Position(cfg.X, cfg.Y, bounds.Dx(), bounds.Dy()), draw.Over)
	}
	nameRes.Sort()

	err := createImage(fmt.Sprintf("%v-%v", baseName, strings.Join(nameRes, "-"))+".png", base)

	return err
}
