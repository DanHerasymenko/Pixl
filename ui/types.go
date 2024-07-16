package ui

import (
	"fyne.io/fyne/v2"
	"zeromastery.io/pixl/apptype"
	"zeromastery.io/pixl/pxcanv"
	"zeromastery.io/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanv.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
