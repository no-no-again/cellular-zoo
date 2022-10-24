package world

import (
	"math/rand"

	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/grid"
	"github.com/zronev/cellular-zoo/rule"
)

type World struct {
	grid *grid.Grid[int]
}

func New(rows, cols, states int) *World {
	grid := grid.New[int](rows, cols)
	cap := int(float64(rows*cols) * config.SpawnCapacity)

	for count := 0; count < cap; count++ {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		grid.Set(c, r, rand.Intn(states))
	}

	return &World{grid}
}

func (c *World) NextGen(rule *rule.Rule) {
	copiedGrid := c.grid.Copy()

	c.grid.Traverse(func(x, y int, cell *int) {
		neighbours := rule.CountNeighbours(x, y, copiedGrid)
		*cell = rule.Apply(*cell, neighbours)
	})
}
