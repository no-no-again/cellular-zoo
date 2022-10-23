package rules

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

// TODO: support different types of neighbourhoods
func (r *Rule) Apply(cell, neighbours int) int {
	switch {
	case cell == 0:
		if r.birth[neighbours] {
			return r.state - 1
		}
		return 0
	case cell == r.state-1:
		if r.survival[neighbours] {
			return r.state - 1
		}
		return r.state - 2
	case cell < r.state-1:
		return cell - 1
	default:
		return cell
	}
}
