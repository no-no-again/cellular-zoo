package drawers

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type WindowDrawer struct {
	imd          *imdraw.IMDraw
	windowWidth  int
	windowHeight int
}

func NewWindowDrawer(imd *imdraw.IMDraw, windowWidth, windowHeight int) *WindowDrawer {
	return &WindowDrawer{imd, windowWidth, windowHeight}
}

func (d *WindowDrawer) DrawRect(x, y, w, h float64, color color.RGBA) {
	y = float64(d.windowHeight) - y

	rect := pixel.R(x, y, x+w, y-h).Norm()

	d.imd.Color = color
	d.imd.Push(rect.Min, rect.Max)
	d.imd.Rectangle(0)
}
