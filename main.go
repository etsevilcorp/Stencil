package main

import (
	"log"
	"os"

	"github.com/etsevilcorp/stencil/stencil"
	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	buf, err := os.ReadFile("./test.toml")
	if err != nil {
		panic(err)
	}

	stencils := make(stencil.Stencils, 1)
	err = toml.Unmarshal(buf, &stencils)
	if err != nil {
		panic(err)
	}

	for name, stencil := range stencils {
		log.Printf("%v: %+v, anchor empty: %v", name, stencil, stencil.Anchor == "")
	}
}
