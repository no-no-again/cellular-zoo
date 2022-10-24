package world

import (
	"github.com/zronev/cellular-zoo/drawers"

	"golang.org/x/image/colornames"
)

type Drawer struct {
	world   *World
	cellSize int
}

func NewDrawer(world *World, cellSize int) *Drawer {
	return &Drawer{world, cellSize}
}

func (cd *Drawer) Draw(drawer drawers.Drawer) {
	cd.world.grid.Traverse(func(x, y int, cell *int) {
		color := colornames.Snow

		switch *cell {
		case 4:
			color = colornames.Cornflowerblue
		case 3:
			color = colornames.Wheat
		case 2:
			color = colornames.Tomato
		case 1:
			color = colornames.Darkslategray
		}

		drawer.DrawRect(
			float64(x*cd.cellSize),
			float64(y*cd.cellSize),
			float64(cd.cellSize),
			float64(cd.cellSize),
			color,
		)
	})
}
