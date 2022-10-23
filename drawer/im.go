package drawer

import (
	"image/color"

	"github.com/zronev/cellular-zoo/config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type IMDrawer struct {
	imd *imdraw.IMDraw
}

func NewIMDrawer(imd *imdraw.IMDraw) *IMDrawer {
	return &IMDrawer{imd}
}

func (d *IMDrawer) DrawRect(x, y, w, h float64, color color.RGBA) {
	y = config.WindowHeight - y

	rect := pixel.R(x, y, x+w, y-h).Norm()

	d.imd.Color = color
	d.imd.Push(rect.Min, rect.Max)
	d.imd.Rectangle(0)
}
