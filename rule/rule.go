package rule

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

// TODO: support different types of neighbourhoods
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
