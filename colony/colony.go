package colony

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/drawers"

	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Cell int

type Colony struct {
	cells []Cell
	rows  int
	cols  int
}

func New(rows, cols int) *Colony {
	cells := make([]Cell, rows*cols)

	c := &Colony{cells, rows, cols}
	c.Traverse(func(x, y int, cell *Cell) {
		*cell = Cell(rand.Intn(2))
	})

	return c
}

func (c *Colony) Get(x, y int) *Cell {
	return &c.cells[c.cols*y+x]
}

func (c *Colony) Set(x, y int, cell Cell) {
	c.cells[c.cols*y+x] = cell
}

func (c *Colony) Traverse(f func(x, y int, cell *Cell)) {
	for row := 0; row < c.rows; row++ {
		for col := 0; col < c.cols; col++ {
			f(col, row, c.Get(col, row))
		}
	}
}

type ColonyDrawer struct {
	colony   *Colony
	cellSize int
}

func NewDrawer(c *Colony, cellSize int) *ColonyDrawer {
	return &ColonyDrawer{c, cellSize}
}

func (cd *ColonyDrawer) Draw(imd *imdraw.IMDraw) {
	cd.colony.Traverse(func(x, y int, cell *Cell) {
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
