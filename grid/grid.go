package grid

import "fmt"

type Grid[T any] struct {
	vals []T
	rows int
	cols int
}

func New[T any](rows, cols int) *Grid[T] {
	vals := make([]T, rows*cols)
	return &Grid[T]{vals, rows, cols}
}

func FromValues[T any](rows, cols int, vals ...T) *Grid[T] {
	return &Grid[T]{vals, rows, cols}
}

func (g *Grid[T]) Rows() int {
	return g.rows
}

func (g *Grid[T]) Cols() int {
	return g.cols
}

func (g *Grid[T]) Get(x, y int) T {
	return g.vals[g.cols*y+x]
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.vals[g.cols*y+x] = v
}

func (g *Grid[T]) Traverse(f func(x, y int, v T)) {
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			f(col, row, g.Get(col, row))
		}
	}
}

func (g *Grid[T]) Copy() *Grid[T] {
	copied := make([]T, g.rows*g.cols)
	copy(copied, g.vals)
	return &Grid[T]{copied, g.rows, g.cols}
}

func (g *Grid[T]) Debug() {
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			fmt.Printf("%v ", g.Get(col, row))
		}
		fmt.Println()
	}
}
