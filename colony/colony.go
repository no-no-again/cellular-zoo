package colony

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/drawers"
	"github.com/zronev/cellular-zoo/grid"

	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Cell int

type Colony struct {
	grid *grid.Grid[Cell]
}

func New(rows, cols int) *Colony {
	grid := grid.New[Cell](rows, cols)

	c := &Colony{grid}
	c.grid.Traverse(func(x, y int, cell *Cell) {
		*cell = Cell(rand.Intn(2))
	})

	return c
}

type ColonyDrawer struct {
	colony   *Colony
	cellSize int
}

func NewDrawer(c *Colony, cellSize int) *ColonyDrawer {
	return &ColonyDrawer{c, cellSize}
}

func (cd *ColonyDrawer) Draw(imd *imdraw.IMDraw) {
	cd.colony.grid.Traverse(func(x, y int, cell *Cell) {
		color := colornames.Tomato

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
