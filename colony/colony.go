package colony

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/grid"
)

type Colony struct {
	grid *grid.Grid[int]
}

func New(rows, cols, states int) *Colony {
	grid := grid.New[int](rows, cols)

	c := &Colony{grid}
	c.grid.Traverse(func(x, y int, cell *int) {
		*cell = int(rand.Intn(states))
	})

	return c
}

func (c *Colony) NextGen(rules func(cell int, neighbours int) int) {
	copiedGrid := c.grid.Copy()

	c.grid.Traverse(func(x, y int, cell *int) {
		neighbours := countNeighbours(x, y, copiedGrid)
		*cell = rules(*cell, neighbours)
	})
}

func countNeighbours(x, y int, g *grid.Grid[int]) int {
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

			if *g.Get(x, y) != 0 {
				neighbours++
			}
		}
	}

	return neighbours
}
