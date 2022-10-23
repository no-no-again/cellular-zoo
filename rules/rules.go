package rules

import "fmt"

type Neighbourhood int

const (
	Moore = Neighbourhood(iota)
	Neumann
)

type Rule struct {
	survival      map[int]bool
	birth         map[int]bool
	state         int
	neighbourhood Neighbourhood
}

func FromString(rawString string) (rule *Rule, err error) {
	return parseRule(rawString)
}

type GOLCell int

func GOL(cell GOLCell, neighbours int) GOLCell {
	switch cell {
	case GOLCell(0):
		if neighbours == 3 {
			return GOLCell(1)
		}
		return GOLCell(0)
	case GOLCell(1):
		if neighbours < 2 || neighbours > 3 {
			return GOLCell(0)
		}
		return GOLCell(1)
	default:
		panic(fmt.Sprintf("unknown cell state: %d", cell))
	}
}
