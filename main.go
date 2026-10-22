package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/etsevilcorp/stencil/engine"
	"github.com/etsevilcorp/stencil/stencil"
	"github.com/etsevilcorp/stencil/tui"
	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	st := flag.String("stencil", "./stencil.toml", "Stencil to use")
	base := flag.String("base", "./base.png", "Base image to attach assets to")
	out := flag.String("out", "./", "Where to output")
	tty := flag.Bool("ui", false, "Use interactive mode")

	flag.Parse()

	buf, err := os.ReadFile(*st)
	if err != nil {
		log.Panic(err.Error())
	}

	stencils := make(stencil.Stencils)
	err = toml.Unmarshal(buf, &stencils)
	if err != nil {
		log.Panic(err.Error())
	}

	err = os.MkdirAll(filepath.Dir(*out), 0755)
	if err != nil {
		log.Panic(err.Error())
	}

	if *tty {
		tui.Start(stencils, *base)
	} else {
		err = engine.Position(stencils, *base, *out)
		if err != nil {
			log.Panic(err.Error())
		}
	}
}
