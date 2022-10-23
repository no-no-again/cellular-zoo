package rule

import "github.com/zronev/cellular-zoo/grid"

type Neighbourhood int

const (
	Moore = Neighbourhood(iota)
	Neumann
)

type Rule struct {
	survival      map[int]bool
	birth         map[int]bool
	states        int
	neighbourhood Neighbourhood
}

func FromString(rawString string) (rule *Rule, err error) {
	return parseRule(rawString)
}

func (r *Rule) States() int {
	return r.states
}

func (r *Rule) Apply(cell, neighbours int) int {
	switch {
	case cell == 0:
		if r.birth[neighbours] {
			return r.states - 1
		}
		return 0
	case cell == r.states-1:
		if r.survival[neighbours] {
			return r.states - 1
		}
		return r.states - 2
	case cell < r.states-1:
		return cell - 1
	default:
		return cell
	}
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
