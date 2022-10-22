package rules

import "testing"

func TestGOL(t *testing.T) {
	cases := []struct {
		neighbours int
		cell       GOLCell
		want       GOLCell
	}{
		{cell: GOLCell(1), neighbours: 2, want: GOLCell(1)},
		{cell: GOLCell(1), neighbours: 3, want: GOLCell(1)},
		{cell: GOLCell(1), neighbours: 1, want: GOLCell(0)},
		{cell: GOLCell(1), neighbours: 4, want: GOLCell(0)},

		{cell: GOLCell(0), neighbours: 3, want: GOLCell(1)},
		{cell: GOLCell(0), neighbours: 2, want: GOLCell(0)},
		{cell: GOLCell(0), neighbours: 4, want: GOLCell(0)},
	}

	for _, tt := range cases {
		got := GOL(tt.cell, tt.neighbours)
		if got != tt.want {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}
