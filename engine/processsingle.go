package engine

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "image/jpeg"
	"image/png"

	"github.com/etsevilcorp/stencil/stencil"
)

func ProcessSingle(base draw.Image, baseName string, stencils map[string]string, cfgs stencil.Stencils) error {

	var nameRes = make(sort.StringSlice, 0, len(stencils))
	for name, path := range stencils {
		cfg := cfgs[name]
		err := cfg.Validate()
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

		draw.Draw(base, base.Bounds(), stenciling, cfg.Anchor.Position(cfg.X, cfg.Y, bounds.Dx(), bounds.Dy()), draw.Src)
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
