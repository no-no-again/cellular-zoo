package colony

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/grid"
	"github.com/zronev/cellular-zoo/rule"
)

type Colony struct {
	grid *grid.Grid[int]
}

func New(rows, cols, states int) *Colony {
	grid := grid.New[int](rows, cols)

	for count := 0; count < (rows*cols)/3.0; count++ {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		grid.Set(c, r, rand.Intn(states))
	}

	return &Colony{grid}
}

func (c *Colony) NextGen(rule *rule.Rule) {
	copiedGrid := c.grid.Copy()

	c.grid.Traverse(func(x, y int, cell *int) {
		neighbours := countNeighbours(x, y, copiedGrid)
		*cell = rule.Apply(*cell, neighbours)
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
