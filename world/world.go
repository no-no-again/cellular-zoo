package world

import (
	"math/rand"
	"runtime"
	"sync"

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
	var wg sync.WaitGroup
	nworkers := runtime.NumCPU()
	gridCopy := c.grid.Copy()

	for i := 0; i < nworkers; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			rowsPerWorker := c.grid.Rows() / nworkers

			for y := i * rowsPerWorker; y < c.grid.Rows() && y < (i+1)*rowsPerWorker; y++ {
				for x := 0; x < c.grid.Cols(); x++ {
					cell := c.grid.Get(x, y)
					neighbours := rule.CountNeighbours(x, y, gridCopy)
					*cell = rule.Apply(*cell, neighbours)
				}
			}
		}(i)
	}

	wg.Wait()
}
