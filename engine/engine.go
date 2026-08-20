package engine

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/etsevilcorp/stencil/stencil"
)

// Position returs error on fail(of any kind).
// stencils are parsed stencil map, baseImage is path to base image
func Position(stencils stencil.Stencils, baseImage string) error {
	// baseFile, err := os.Open(baseImage)
	// if err != nil {
	// 	return fmt.Errorf("position: failed to load an image %v: %w", baseImage, err)
	// }
	// defer baseFile.Close()

	// base, name, err := image.Decode(baseFile)
	// if err != nil {
	// 	return fmt.Errorf("position: failed to convert file to an image %v: %w", baseImage, err)
	// }
	// baseDraw := ConvertToDrawImage(base)

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
	panic(fmt.Sprintf("%+v", combos))
}
