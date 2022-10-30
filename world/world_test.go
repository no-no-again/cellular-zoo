package world

import (
	"reflect"
	"testing"

	"github.com/zronev/cellular-zoo/grid"
	"github.com/zronev/cellular-zoo/rule"
)

func TestWorld(t *testing.T) {
	allDeadAfterFirstStepRule, _ := rule.FromString("0/0/2/M")

	cases := []struct {
		nworkers int
		rule     *rule.Rule
		grid     *grid.Grid[int]
		want     *grid.Grid[int]
	}{
		{
			nworkers: 5,
			rule:     allDeadAfterFirstStepRule,
			grid: grid.FromValues(
				7, 4,
				0, 1, 0, 1,
				0, 1, 0, 1,
				0, 1, 0, 1,
				0, 1, 0, 1,
				0, 1, 0, 1,
				0, 1, 0, 1,
				0, 1, 0, 1,
			),
			want: grid.FromValues(
				7, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			),
		},
	}

	for _, tt := range cases {
		w := FromGrid(tt.grid)
		w.nextGen(tt.rule, tt.nworkers)

		got := w.grid

		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("want: %v, got: %v", *tt.want, *got)
		}
	}
}
