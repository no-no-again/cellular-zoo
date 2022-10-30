package rule

import (
	"testing"

	"github.com/zronev/cellular-zoo/grid"
)

func TestGOLRule(t *testing.T) {
	rule, _ := FromString("2-3/3/2/M")

	cases := []struct {
		neighbours int
		cell       int
		want       int
	}{
		{cell: 1, neighbours: 2, want: 1},
		{cell: 1, neighbours: 3, want: 1},
		{cell: 1, neighbours: 1, want: 0},
		{cell: 1, neighbours: 4, want: 0},

		{cell: 0, neighbours: 3, want: 1},
		{cell: 0, neighbours: 2, want: 0},
		{cell: 0, neighbours: 4, want: 0},
	}

	for _, tt := range cases {
		got := rule.Apply(tt.cell, tt.neighbours)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}

func TestRuleWithManyStates(t *testing.T) {
	rule, _ := FromString("2-3/3/3/M")

	cases := []struct {
		neighbours int
		cell       int
		want       int
	}{
		{cell: 2, neighbours: 2, want: 2},
		{cell: 2, neighbours: 3, want: 2},

		{cell: 1, neighbours: 2, want: 0},
		{cell: 1, neighbours: 3, want: 0},

		{cell: 2, neighbours: 1, want: 1},
		{cell: 2, neighbours: 4, want: 1},

		{cell: 0, neighbours: 3, want: 2},
		{cell: 0, neighbours: 2, want: 0},
		{cell: 0, neighbours: 4, want: 0},
	}

	for _, tt := range cases {
		got := rule.Apply(tt.cell, tt.neighbours)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}

func TestCountNeighbours(t *testing.T) {
	cases := []struct {
		grid *grid.Grid[int]
		rule *Rule
		x, y int
		want int
	}{
		{
			want: 0,
			x:    1,
			y:    1,
			rule: mustGetRule("2/2/2/M"),
			grid: grid.FromValues(
				3, 3,
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			),
		},
		{
			want: 5,
			x:    1,
			y:    1,
			rule: mustGetRule("2/2/2/M"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},
		{
			want: 2,
			x:    0,
			y:    0,
			rule: mustGetRule("2/2/2/M"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},
		{
			want: 3,
			x:    2,
			y:    2,
			rule: mustGetRule("2/2/2/M"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},

		{
			want: 0,
			x:    1,
			y:    1,
			rule: mustGetRule("2/2/2/N"),
			grid: grid.FromValues(
				3, 3,
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			),
		},
		{
			want: 3,
			x:    1,
			y:    1,
			rule: mustGetRule("2/2/2/N"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},
		{
			want: 1,
			x:    0,
			y:    0,
			rule: mustGetRule("2/2/2/N"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},
		{
			want: 2,
			x:    2,
			y:    2,
			rule: mustGetRule("2/2/2/N"),
			grid: grid.FromValues(
				3, 3,
				0, 1, 1,
				0, 1, 1,
				1, 1, 0,
			),
		},
	}

	for _, tt := range cases {
		got := tt.rule.CountNeighbours(tt.x, tt.y, tt.grid)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}

func mustGetRule(ruleString string) *Rule {
	rule, _ := FromString(ruleString)
	return rule
}
