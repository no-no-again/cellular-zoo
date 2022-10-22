package drawers

import (
	"image/color"

	"github.com/zronev/cellular-zoo/config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func DrawRect(imd *imdraw.IMDraw, x, y, w, h float64, color color.RGBA) {
	y = config.WindowHeight - y

	rect := pixel.R(x, y, x+w, y-h).Norm()

	imd.Color = color
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
}
