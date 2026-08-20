package main

import (
	"os"

	"github.com/etsevilcorp/stencil/engine"
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

	err = engine.Position(stencils, "./base.png")
	if err != nil {
		panic(err)
	}
}
