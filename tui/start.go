package tui

import (
	"cmp"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"slices"

	"github.com/awesome-gocui/gocui"
	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/etsevilcorp/stencil/engine"
	"github.com/etsevilcorp/stencil/imageutils"
	"github.com/etsevilcorp/stencil/stencil"
	"github.com/zeebo/xxh3"
)

func Start(stencils stencil.Stencils, basePath string) error {
	g, err := gocui.NewGui(gocui.OutputTrue, true) // 24-bit colour
	if err != nil {
		return err
	}
	defer g.Close()

	base, err := imageutils.OpenImage(basePath)
	if err != nil {
		return fmt.Errorf("position: failed to load an image %v: %w", basePath, err)
	}

	g.SetManagerFunc(layout(base, stencils))

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}
	return nil
}

func insertColoredLayer(s []engine.Layer, count int, st stencil.Stencil, name string) int {
	i, _ := slices.BinarySearchFunc(s[:count], st.ZIndex, func(l engine.Layer, target int) int {
		return cmp.Compare(l.Cfg.ZIndex, target)
	})
	copy(s[i+1:count+1], s[i:count])

	hash := xxh3.HashString(name)
	img := engine.Solid(
		color.RGBA{
			R: uint8((hash >> 16) & 0xFF),
			G: uint8((hash >> 8) & 0xFF),
			B: uint8(hash & 0xFF),
			A: 0xFF,
		}, *st.MaxHeight, *st.MaxHeight)

	s[i] = engine.Layer{
		Name:  name,
		Cfg:   st,
		Image: img,
	}
	return count + 1
}

func layout(img image.Image, stencils stencil.Stencils) func(*gocui.Gui) error {
	var lastW, lastH int // size we last rendered at
	base := engine.Solid(color.White, img.Bounds().Dx(), img.Bounds().Dy())

	layers := make([]engine.Layer, len(stencils))
	i := 0
	for name, stencil := range stencils {
		i = insertColoredLayer(layers, i, stencil, name)
	}
	log.SetOutput(os.Stderr)
	log.Printf("%+v", layers)

	engine.Draw(base, layers)

	return func(g *gocui.Gui) error {
		maxX, maxY := g.Size()

		v, err := g.SetView("test", 0, 0, maxX-1, maxY-1, 0)
		if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}

		// Usable area inside the frame.
		w, h := maxX-2, maxY-2
		if w <= 0 || h <= 0 || (w == lastW && h == lastH) {
			return nil
		}
		lastW, lastH = w, h

		// Each terminal row holds two vertical pixels (the half-block trick),
		// so the pixel height is h*2.
		ans, err := ansimage.NewScaledFromImage(
			base, h*2, w,
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			ansimage.ScaleModeFit,
			ansimage.NoDithering,
		)

		if err != nil {
			return err
		}

		v.Clear()
		fmt.Fprint(v, ans.Render())
		return nil
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
