package pkg

import "math/rand"

// City encapsulates the neighbour informations of the city
type City struct {
	Name  string
	West  *City
	East  *City
	North *City
	South *City
}

// NewCity returns a new instance of the city
func NewCity(name string) *City {
	return &City{
		Name: name,
	}
}

// GetRandomNeighbour returns one of the neighbours in random,
// returns nil if there is no neighbours
func (c *City) GetRandomNeighbour() *City {
	availableNeighbours := []*City{}
	if c.West != nil {
		availableNeighbours = append(availableNeighbours, c.West)
	}
	if c.East != nil {
		availableNeighbours = append(availableNeighbours, c.East)
	}
	if c.North != nil {
		availableNeighbours = append(availableNeighbours, c.North)
	}
	if c.South != nil {
		availableNeighbours = append(availableNeighbours, c.South)
	}
	if len(availableNeighbours) == 0 {
		return nil
	}
	return availableNeighbours[rand.Int()%(len(availableNeighbours))]
}
