package colony

import (
	"github.com/zronev/cellular-zoo/drawers"

	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type ColonyDrawer struct {
	colony   *Colony
	cellSize int
}

func NewDrawer(c *Colony, cellSize int) *ColonyDrawer {
	return &ColonyDrawer{c, cellSize}
}

func (cd *ColonyDrawer) Draw(imd *imdraw.IMDraw) {
	cd.colony.grid.Traverse(func(x, y int, cell *Cell) {
		color := colornames.Black

		if *cell == Cell(0) {
			color = colornames.Snow
		}

		drawers.DrawRect(
			imd,
			float64(x*cd.cellSize),
			float64(y*cd.cellSize),
			float64(cd.cellSize),
			float64(cd.cellSize),
			color,
		)
	})
}
