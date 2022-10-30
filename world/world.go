package world

import (
	"math/rand"
	"runtime"
	"sync"

	"github.com/no-no-again/cellular-zoo/grid"
	"github.com/no-no-again/cellular-zoo/rule"
)

type World struct {
	grid *grid.Grid[int]
}

func New(rows, cols int, rule *rule.Rule, spawnCap float64) *World {
	w := &World{
		grid: grid.New[int](rows, cols),
	}
	w.Spawn(rule, spawnCap)
	return w
}

func FromGrid(grid *grid.Grid[int]) *World {
	return &World{grid}
}

func (w *World) NextGen(rule *rule.Rule) {
	nworkers := runtime.NumCPU()
	w.nextGen(rule, nworkers)
}

// TODO: check how bad it is to spawn a bunch of goroutines on every frame.
// Maybe rewrite this using long-living workers and channels.
// But this is good enough for now.
func (w *World) nextGen(rule *rule.Rule, nworkers int) {
	var wg sync.WaitGroup

	gridCopy := w.grid.Copy()
	rowsPerWorker := w.grid.Rows()/nworkers + 1

	for i := 0; i < nworkers; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for y := i * rowsPerWorker; y < w.grid.Rows() && y < (i+1)*rowsPerWorker; y++ {
				for x := 0; x < w.grid.Cols(); x++ {
					cell := w.grid.Get(x, y)
					neighbours := rule.CountNeighbours(x, y, gridCopy)

					newCell := rule.Apply(cell, neighbours)
					w.grid.Set(x, y, newCell)
				}
			}
		}(i)
	}

	wg.Wait()
}

func (w *World) Spawn(rule *rule.Rule, spawnCap float64) {
	rows := w.grid.Rows()
	cols := w.grid.Cols()

	cap := int(float64(rows*cols) * spawnCap)

	for count := 0; count < cap; count++ {
		x := rand.Intn(cols)
		y := rand.Intn(rows)
		states := rand.Intn(rule.States())
		w.grid.Set(x, y, states)
	}
}

func (w *World) Clear() {
	w.grid = grid.New[int](w.grid.Rows(), w.grid.Cols())
}
