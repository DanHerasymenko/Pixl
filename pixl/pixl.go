package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"image/color"
	"zeromastery.io/pixl/apptype"
	"zeromastery.io/pixl/pxcanv"
	"zeromastery.io/pixl/swatch"
	"zeromastery.io/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor: color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.Size{600, 600},
		CanvasOffset: fyne.Position{0, 0},
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}
	pixlCanvas := pxcanv.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()

}
