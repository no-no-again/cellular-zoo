package drawers

import (
	"image/color"
)

type Drawer interface {
	DrawRect(x, y, w, h float64, color color.RGBA)
}
