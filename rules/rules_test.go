package rules

import "testing"

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
