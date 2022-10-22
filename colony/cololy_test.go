package colony

import (
	"testing"

	"github.com/zronev/cellular-zoo/grid"
)

func TestCountNeighbours(t *testing.T) {
	cases := []struct {
		grid *grid.Grid[Cell]
		want int
		x, y int
	}{
		{
			want: 0,
			x:    1,
			y:    1,
			grid: grid.NewFromValues(
				3, 3,
				Cell(0), Cell(0), Cell(0),
				Cell(0), Cell(0), Cell(0),
				Cell(0), Cell(0), Cell(0),
			),
		},
		{
			want: 5,
			x:    1,
			y:    1,
			grid: grid.NewFromValues(
				3, 3,
				Cell(0), Cell(1), Cell(1),
				Cell(0), Cell(1), Cell(1),
				Cell(1), Cell(1), Cell(0),
			),
		},
		{
			want: 2,
			x:    0,
			y:    0,
			grid: grid.NewFromValues(
				3, 3,
				Cell(0), Cell(1), Cell(1),
				Cell(0), Cell(1), Cell(1),
				Cell(1), Cell(1), Cell(0),
			),
		},
		{
			want: 3,
			x:    2,
			y:    2,
			grid: grid.NewFromValues(
				3, 3,
				Cell(0), Cell(1), Cell(1),
				Cell(0), Cell(1), Cell(1),
				Cell(1), Cell(1), Cell(0),
			),
		},
	}

	for _, tt := range cases {
		got := countNeighbours(tt.x, tt.y, tt.grid)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}

func TestGOLRules(t *testing.T) {
	cases := []struct {
		neighbours int
		cell       Cell
		want       Cell
	}{
		{cell: Cell(1), neighbours: 2, want: Cell(1)},
		{cell: Cell(1), neighbours: 3, want: Cell(1)},
		{cell: Cell(1), neighbours: 1, want: Cell(0)},
		{cell: Cell(1), neighbours: 4, want: Cell(0)},

		{cell: Cell(0), neighbours: 3, want: Cell(1)},
		{cell: Cell(0), neighbours: 2, want: Cell(0)},
		{cell: Cell(0), neighbours: 4, want: Cell(0)},
	}

	for _, tt := range cases {
		got := golRules(tt.cell, tt.neighbours)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}
