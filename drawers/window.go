package drawers

import (
	"image/color"

	"github.com/zronev/cellular-zoo/config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type WindowDrawer struct {
	imd *imdraw.IMDraw
}

func NewWindowDrawer(imd *imdraw.IMDraw) *WindowDrawer {
	return &WindowDrawer{imd}
}

func (d *WindowDrawer) DrawRect(x, y, w, h float64, color color.RGBA) {
	y = config.WindowHeight - y

	rect := pixel.R(x, y, x+w, y-h).Norm()

	d.imd.Color = color
	d.imd.Push(rect.Min, rect.Max)
	d.imd.Rectangle(0)
}
