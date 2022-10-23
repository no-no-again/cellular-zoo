package colony

import (
	"github.com/zronev/cellular-zoo/drawer"

	"golang.org/x/image/colornames"
)

type ColonyDrawer struct {
	drawer   drawer.Drawer
	colony   *Colony
	cellSize int
}

func NewDrawer(drawer drawer.Drawer, colony *Colony, cellSize int) *ColonyDrawer {
	return &ColonyDrawer{drawer, colony, cellSize}
}

func (cd *ColonyDrawer) Draw() {
	cd.colony.grid.Traverse(func(x, y int, cell *int) {
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

		cd.drawer.DrawRect(
			float64(x*cd.cellSize),
			float64(y*cd.cellSize),
			float64(cd.cellSize),
			float64(cd.cellSize),
			color,
		)
	})
}
