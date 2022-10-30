package world

import (
	"image/color"

	"github.com/zronev/cellular-zoo/drawers"
)

type Palette []color.RGBA

type Drawer struct {
	world    *World
	palette  Palette
	cellSize int
}

func NewDrawer(world *World, palette Palette, cellSize int) *Drawer {
	return &Drawer{world, palette, cellSize}
}

func (wd *Drawer) Draw(drawer drawers.Drawer) {
	wd.world.grid.Traverse(func(x, y int, cell *int) {

		drawer.DrawRect(
			float64(x*wd.cellSize),
			float64(y*wd.cellSize),
			float64(wd.cellSize),
			float64(wd.cellSize),
			wd.palette[*cell],
		)
	})
}
