package engine

import (
	"slices"

	"github.com/etsevilcorp/stencil/stencil"
)

func sortNames(stencils map[string]string, cfgs stencil.Stencils) []string {
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

	return stencilsSorted
}
