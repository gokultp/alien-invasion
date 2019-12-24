package pkg

import (
	"reflect"
	"sync"
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

func TestCity_Isolate(t *testing.T) {
	sf := NewCity("sf")
	berlin := NewCity("newyork")
	tokyo := NewCity("tokyo")
	tokyo.West = berlin
	berlin.East = tokyo
	sf.East = berlin
	berlin.West = sf
	banglore := NewCity("banglore")

	tests := []struct {
		name string
		city *City
	}{
		{
			name: "case of an already isolated city",
			city: banglore,
		},
		{
			name: "checking if isolating a connected city",
			city: berlin,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.city.Isolate()
			if tt.city.West != nil && tt.city.West.East != nil {
				t.Error("error: not Isolated")
			}
			if tt.city.East != nil && tt.city.East.West != nil {
				t.Error("error: not Isolated")
			}
			if tt.city.South != nil && tt.city.South.North != nil {
				t.Error("error: not Isolated")
			}
			if tt.city.North != nil && tt.city.North.South != nil {
				t.Error("error: not Isolated")
			}
		})
	}
}

func TestCity_Aquire(t *testing.T) {
	alien1 := 1
	tests := []struct {
		name string
		c    *City
		args int
		want int
	}{
		{
			name: "should return the same alien id if city is free",
			c:    NewCity("sf"),
			args: 1,
			want: 1,
		},
		{
			name: "should return the the aquired alien's id if already aquired.",
			c:    &City{alienAquired: &alien1, lock: &sync.Mutex{}},
			args: 2,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.c.Aquire(tt.args); got != tt.want {
				t.Errorf("City.Aquire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCity_Free(t *testing.T) {
	alien1 := 1

	tests := []struct {
		name string
		c    *City
	}{
		{
			name: "should get free if invoked",
			c: &City{
				alienAquired: &alien1,
				lock:         &sync.Mutex{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Free()
			if tt.c.alienAquired != nil {
				t.Error("city did not get free")
			}
		})
	}
}
