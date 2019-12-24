package pkg

import (
	"fmt"
	"math/rand"
	"sync"
)

// City encapsulates the neighbour informations of the city
type City struct {
	Name         string
	West         *City
	East         *City
	North        *City
	South        *City
	alienAquired *int
	lock         *sync.Mutex
}

// NewCity returns a new instance of the city
func NewCity(name string) *City {
	return &City{
		Name: name,
		lock: &sync.Mutex{},
	}
}

// String implemets stringer interface
func (c *City) String() string {
	str := c.Name
	if c.West != nil {
		str = fmt.Sprintf("%s west=%s", str, c.West.Name)
	}
	if c.East != nil {
		str = fmt.Sprintf("%s east=%s", str, c.East.Name)
	}
	if c.North != nil {
		str = fmt.Sprintf("%s north=%s", str, c.North.Name)
	}
	if c.South != nil {
		str = fmt.Sprintf("%s south=%s", str, c.South.Name)
	}
	return str
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
	return availableNeighbours[rand.Intn(len(availableNeighbours))]
}

// Isolate will remove connections from neighbouring cities to the given city
func (c *City) Isolate() {
	if c.West != nil {
		c.West.East = nil
	}
	if c.East != nil {
		c.East.West = nil
	}
	if c.North != nil {
		c.North.South = nil
	}
	if c.South != nil {
		c.South.North = nil
	}
}

// Aquire will aquire the city and assign to the given alien if it is free
// returns the id of alien aquired it first
func (c *City) Aquire(alien int) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.alienAquired == nil {
		c.alienAquired = &alien
		return alien
	}
	return *c.alienAquired
}

// Free will unset the aquisition in the city
func (c *City) Free() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.alienAquired = nil
}

// IsIsolated says wheher the city is isolated or not
// return true if it does not have any neighbouring cities
// return false otherwse.
func (c *City) IsIsolated() bool {
	return c.West == nil && c.East == nil && c.North == nil && c.South == nil
}
