package engine

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	_ "image/jpeg"
	"image/png"

	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils) error {
	cfgsSorted := make([]stencil.Stencil, 0, len(cfgs))
	stencilsSorted := make([]string, 0, len(stencils))
	for name, cfg := range cfgs {
		i, ok := slices.BinarySearchFunc(cfgsSorted, cfg, func(el, tar stencil.Stencil) int {
			return el.ZIndex - tar.ZIndex
		})
		if !ok {
			i = len(cfgsSorted)
		}
		cfgsSorted = slices.Insert(cfgsSorted, i, cfg)
		stencilsSorted = slices.Insert(stencilsSorted, i, name)
	}

	log.Printf("%+v", cfgsSorted)

	var nameRes = make(sort.StringSlice, 0, len(stencils))
	for _, name := range stencilsSorted {
		cfg := cfgs[name]
		err := cfg.Validate()

		path := stencils[name]
		if err != nil {
			return err
		}

		stFile, err := os.Open(filepath.Join(name, path))
		if err != nil {
			return err
		}
		nameRes = append(nameRes, (strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))))

		stenciling, _, err := image.Decode(stFile)
		if err != nil {
			return err
		}

		bounds := stenciling.Bounds()
		if cfg.MaxHeight != nil && *cfg.MaxHeight < bounds.Max.Y || cfg.MaxWidth != nil && *cfg.MaxWidth < bounds.Max.X {
			stenciling = cfg.HandleMax(stenciling)
		}
		if cfg.MinHeight != nil && *cfg.MinHeight > bounds.Max.Y || cfg.MinWidth != nil && *cfg.MinWidth > bounds.Max.X {
			stenciling = cfg.HandleMin(stenciling)
		}

		draw.Draw(base, base.Bounds(), stenciling, cfg.Anchor.Position(cfg.X, cfg.Y, bounds.Dx(), bounds.Dy()), draw.Over)
	}
	nameRes.Sort()

	outputFile, err := os.Create(fmt.Sprintf("%v-%v", baseName, strings.Join(nameRes, "-")) + ".png")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, base)
	if err != nil {
		return err
	}

	return nil
}
