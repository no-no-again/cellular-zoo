package colony

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/grid"
)

const (
	numberOfStates = 2
)

type Cell int

type Colony struct {
	grid *grid.Grid[Cell]
}

func New(rows, cols int) *Colony {
	grid := grid.New[Cell](rows, cols)

	c := &Colony{grid}
	c.grid.Traverse(func(x, y int, cell *Cell) {
		*cell = Cell(rand.Intn(numberOfStates))
	})

	return c
}

func (c *Colony) NextGen() {
	c.grid.Traverse(func(x, y int, cell *Cell) {
		*cell = Cell(rand.Intn(numberOfStates))
	})
}
