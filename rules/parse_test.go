package rules

import (
	"reflect"
	"testing"
)

func TestParseRule(t *testing.T) {
	cases := []struct {
		rawRule string
		want    *Rule
	}{
		{
			rawRule: "1/2/2/M",
			want: &Rule{
				survival: map[int]bool{
					1: true,
				},
				birth: map[int]bool{
					2: true,
				},
				states:        2,
				neighbourhood: Moore,
			},
		},
		{
			rawRule: "1/2/2/N",
			want: &Rule{
				survival: map[int]bool{
					1: true,
				},
				birth: map[int]bool{
					2: true,
				},
				states:        2,
				neighbourhood: Neumann,
			},
		},
		{
			rawRule: "1,2/2,3/2/N",
			want: &Rule{
				survival: map[int]bool{
					1: true,
					2: true,
				},
				birth: map[int]bool{
					2: true,
					3: true,
				},
				states:        2,
				neighbourhood: Neumann,
			},
		},
		{
			rawRule: "1-2,3/2,3-4/2/N",
			want: &Rule{
				survival: map[int]bool{
					1: true,
					2: true,
					3: true,
				},
				birth: map[int]bool{
					2: true,
					3: true,
					4: true,
				},
				states:        2,
				neighbourhood: Neumann,
			},
		},
	}

	for _, tt := range cases {
		// TODO: check for errors
		got, _ := parseRule(tt.rawRule)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("want: %v, got: %v", *tt.want, *got)
		}
	}
}
