package colony

import (
	"fmt"
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
	copiedGrid := c.grid.Copy()

	c.grid.Traverse(func(x, y int, cell *Cell) {
		neighbours := countNeighbours(x, y, copiedGrid)
		state := golRules(*cell, neighbours)
		*cell = state
	})
}

func countNeighbours(x, y int, g *grid.Grid[Cell]) int {
	neighbours := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			x := x + i
			y := y + j

			if (x < 0 || x >= g.Cols()) ||
				(y < 0 || y >= g.Rows()) {
				continue
			}

			if *g.Get(x, y) == Cell(1) {
				neighbours++
			}
		}
	}

	return neighbours
}

func golRules(cell Cell, neighbours int) Cell {
	switch cell {
	case Cell(0):
		if neighbours == 3 {
			return Cell(1)
		}
		return Cell(0)
	case Cell(1):
		if neighbours < 2 || neighbours > 3 {
			return Cell(0)
		}
		return Cell(1)
	default:
		panic(fmt.Sprintf("unknown cell state: %d", cell))
	}
}
