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

	for count := 0; count < (rows*cols)/4.0; count++ {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		grid.Set(c, r, rand.Intn(states))
	}

	return &Colony{grid}
}

func (c *Colony) NextGen(rule *rule.Rule) {
	copiedGrid := c.grid.Copy()

	c.grid.Traverse(func(x, y int, cell *int) {
		neighbours := rule.CountNeighbours(x, y, copiedGrid)
		*cell = rule.Apply(*cell, neighbours)
	})
}
