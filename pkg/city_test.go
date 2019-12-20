package pkg

import (
	"reflect"
	"testing"
)

func TestCity_getRandomNeighbour(t *testing.T) {
	sf := NewCity("sf")
	berlin := NewCity("newyork")
	tokyo := NewCity("tokyo")
	tokyo.West = berlin
	berlin.East = tokyo
	sf.East = berlin
	berlin.West = sf
	banglore := NewCity("banglore")

	tests := []struct {
		name  string
		input *City
		want  []*City
	}{
		{
			name:  "should return null if there is no neighbours",
			input: banglore,
			want:  nil,
		},
		{
			name:  "should return the neighbour if there is only one neighbour",
			input: sf,
			want:  []*City{berlin},
		},
		{
			name:  "should return any of the neighbours if there is more than one neighbour",
			input: berlin,
			want:  []*City{sf, tokyo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.GetRandomNeighbour()
			if tt.want == nil && got == nil {
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("exptected <null>, got %v", got)
			}
			matched := false
			for _, n := range tt.want {
				if reflect.DeepEqual(n, got) {
					matched = true
					break
				}
			}
			if !matched {
				t.Errorf("got %v, expected one of %v", tt.want, tt.want)
			}
		})
	}
}
