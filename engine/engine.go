package engine

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/etsevilcorp/stencil/stencil"
)

// Position returs error on fail(of any kind).
// stencils are parsed stencil map, basePath is path to base image
func Position(stencils stencil.Stencils, basePath string) error {
	combosCount := 1
	entriesMatrix := make([]struct {
		entries []os.DirEntry
		label   string
	}, len(stencils))

	var err error
	i := 0
	for stencilsDirPath, _ := range stencils {
		entriesMatrix[i].entries, err = os.ReadDir(stencilsDirPath)
		if err != nil {
			return err
		}
		entriesMatrix[i].label = stencilsDirPath

		combosCount *= len(entriesMatrix[i].entries)

		i++
	}

	log.Printf("%+v", entriesMatrix)

	// []map[name_of_the_element]path
	combos := make([]map[string]string, combosCount)

	combinations(entriesMatrix, 0, map[string]string{}, combos, 0)

	baseFile, err := os.Open(basePath)
	if err != nil {
		return fmt.Errorf("position: failed to load an image %v: %w", basePath, err)
	}
	defer baseFile.Close()

	base, _, err := image.Decode(baseFile)
	if err != nil {
		return fmt.Errorf("position: failed to convert file to an image %v: %w", basePath, err)
	}
	baseDraw := ConvertToRGBA(base)

	bName := strings.TrimSuffix(filepath.Base(basePath), filepath.Ext(basePath))
	for _, combo := range combos {
		b := CopyRGBA(baseDraw)
		err := ProcessSingle(b, bName, combo, stencils)
		if err != nil {
			return err
		}
	}

	return nil
}
