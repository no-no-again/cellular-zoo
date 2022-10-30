package rule

import (
	"fmt"

	"github.com/zronev/cellular-zoo/grid"
)

const (
	Moore = iota
	Neumann
)

type Rule struct {
	survival      map[int]bool
	birth         map[int]bool
	states        int
	neighbourhood int
}

func FromString(rawString string) (rule *Rule, err error) {
	return parseRule(rawString)
}

func (r *Rule) States() int {
	return r.states
}

func (r *Rule) Apply(cell, neighbours int) int {
	switch {
	case r.IsCellDead(cell):
		if r.birth[neighbours] {
			return r.LiveCell()
		}
		return r.DeadCell()
	case r.IsCellAlive(cell):
		if r.survival[neighbours] {
			return r.LiveCell()
		}
		return r.DyingCell(cell)
	case r.IsCellDying(cell):
		return r.DyingCell(cell)
	default:
		panic(fmt.Sprintf("a cell with an impossible state: %d\n", cell))
	}
}

func (r *Rule) DeadCell() int {
	return 0
}

func (r *Rule) LiveCell() int {
	return r.states - 1
}

func (r *Rule) DyingCell(cell int) int {
	return cell - 1
}

func (r *Rule) IsCellDead(cell int) bool {
	return cell == 0
}

func (r *Rule) IsCellAlive(cell int) bool {
	return cell == r.states-1
}

func (r *Rule) IsCellDying(cell int) bool {
	return cell > 0 && cell < r.states-1
}

func (r *Rule) CountNeighbours(x, y int, g *grid.Grid[int]) int {
	neighbours := 0

	switch r.neighbourhood {
	case Moore:
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
	case Neumann:
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if (i == 0 && j != 0) || (j == 0 && i != 0) {
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
		}
	}

	return neighbours
}
