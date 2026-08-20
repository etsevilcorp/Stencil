package engine

import (
	"maps"
	"os"
)

func combinations(options []struct {
	entries []os.DirEntry
	label   string
}, rowIndex int, current map[string]string, result []map[string]string, accum int) {
	if rowIndex == len(options) {
		result[accum] = maps.Clone(current)
		return
	}

	for i := 0; i < len(options[rowIndex].entries); i++ {
		current[options[rowIndex].label] = options[rowIndex].entries[i].Name()

		combinations(options, rowIndex+1, current, result, accum+i+rowIndex)

		delete(current, options[rowIndex].label)
	}
}
