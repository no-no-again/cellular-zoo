package colony

import (
	"github.com/zronev/cellular-zoo/drawer"

	"golang.org/x/image/colornames"
)

type Drawer struct {
	colony   *Colony
	cellSize int
}

func NewDrawer(colony *Colony, cellSize int) *Drawer {
	return &Drawer{colony, cellSize}
}

func (cd *Drawer) Draw(drawer drawer.Drawer) {
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

		drawer.DrawRect(
			float64(x*cd.cellSize),
			float64(y*cd.cellSize),
			float64(cd.cellSize),
			float64(cd.cellSize),
			color,
		)
	})
}
