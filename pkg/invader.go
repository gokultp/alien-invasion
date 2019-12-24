package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
)

// Invader simulates the alien invasion.
type Invader struct {
	Cities         map[string]*City
	AliensPosition map[int]*City
}

// NewInvader creates a new instance and add cities by parsing given Reader interface.
// It can be stdin, a network reader or a file
func NewInvader(r io.Reader) *Invader {
	cities := map[string]*City{}
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		data := strings.Split(
			strings.TrimSpace(sc.Text()),
			" ",
		)
		city, ok := cities[data[0]]
		if !ok {
			city = NewCity(data[0])
			cities[city.Name] = city
		}
		if len(data) == 1 || data[1] == "" {
			continue
		}
		for _, neighbour := range data[1:] {
			neighbourData := strings.Split(neighbour, "=")
			nc, ok := cities[neighbourData[1]]
			if !ok {
				nc = NewCity(neighbourData[1])
				cities[nc.Name] = nc
			}
			switch neighbourData[0] {
			case "west":
				city.West = nc
				nc.East = city
			case "east":
				city.East = nc
				nc.West = city
			case "north":
				city.North = nc
				nc.South = city
			case "south":
				city.South = nc
				nc.North = city
			}
		}
	}

	return &Invader{
		Cities: cities,
	}
}

// InitAliens will initialise aliens and assign a city to each alien
// This is a randomn allocation but there will be only one alien in one city if len(city) > len(aliegns)
// otherwise more than one alien will be allocated to the city and that city will be destroyed instantly
func (i *Invader) InitAliens(count int) {
	i.AliensPosition = map[int]*City{}
	cityNames := i.getCityNames()
	rand := newRandomiser(cityNames)

	for j := 0; j < count; j++ {
		c := rand.getRandomElement()
		selectedCity := i.Cities[c]
		if val := selectedCity.Aquire(j); val != j {
			i.destroyCity(selectedCity, j, val)
			rand.removeElement(c)
			continue
		}
		i.AliensPosition[j] = selectedCity
	}
}

// destroyCity will remove the links to city from its neighbouring cities and delete
func (i *Invader) destroyCity(c *City, n, m int) {
	c.Isolate()
	fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", c.Name, n, m)
	// remove city
	delete(i.Cities, c.Name)
	// remove aliens
	delete(i.AliensPosition, n)
	delete(i.AliensPosition, m)
}

// getCityNames will return names of cities
func (i *Invader) getCityNames() []string {
	cityNames := []string{}
	for city := range i.Cities {
		cityNames = append(cityNames, city)
	}
	return cityNames
}

// Exec will execute the simulation
func (i *Invader) Exec() {

	for j := 0; j < 10000; j++ {
		if len(i.AliensPosition) == 0 || len(i.Cities) == 0 {
			break
		}
		wg := sync.WaitGroup{}
		for a, c := range i.AliensPosition {
			wg.Add(1)
			go func(a int, c *City, wg *sync.WaitGroup) {
				defer wg.Done()
				next := c.GetRandomNeighbour()
				if next != nil {
					c.Free()
					if val := next.Aquire(a); val != a {
						i.destroyCity(next, a, val)
						return
					}
					i.AliensPosition[a] = next
				}

			}(a, c, &wg)
			wg.Wait()
		}
	}
}

// PrintCities will print the cities in the format to given io.Wirter
// (can be a file, network stream or stdout etc)
func (i *Invader) PrintCities(w io.Writer) {
	visited := map[string]struct{}{}
	for _, city := range i.Cities {
		if _, ok := visited[city.Name]; !ok && !city.IsIsolated() {
			fmt.Fprintln(w, city)
			visited[city.Name] = struct{}{}
		}
	}
}
