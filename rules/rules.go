package rules

import "fmt"

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
