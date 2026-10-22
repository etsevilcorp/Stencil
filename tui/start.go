package tui

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/awesome-gocui/gocui"
	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/etsevilcorp/stencil/stencil"
)

func Start(stencils stencil.Stencils, basePath string) error {
	g, err := gocui.NewGui(gocui.OutputTrue, true) // 24-bit colour
	if err != nil {
		return err
	}
	defer g.Close()

	baseFile, err := os.Open(basePath)
	if err != nil {
		return fmt.Errorf("position: failed to load an image %v: %w", basePath, err)
	}
	defer baseFile.Close()

	base, _, err := image.Decode(baseFile)
	if err != nil {
		return fmt.Errorf("position: failed to convert file to an image %v: %w", basePath, err)
	}

	g.SetManagerFunc(layout(base))

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}
	return nil
}

func layout(img image.Image) func(*gocui.Gui) error {
	var lastW, lastH int // size we last rendered at

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
			img, h*2, w,
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
