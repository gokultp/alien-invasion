package pkg

import (
	"math/rand"
)

type randomiser struct {
	l       int
	maxRand int
	items   []string
}

// newRandomiser returns a new instance of randomiser
func newRandomiser(items []string) *randomiser {
	l := len(items)
	return &randomiser{
		items:   items,
		l:       l,
		maxRand: l,
	}
}

// getRandomElement returns a random element from the list and
// ensures that it does not repeat until all elements got considered.
func (r *randomiser) getRandomElement() string {
	rnd := rand.Intn(r.maxRand)
	t := r.items[rnd]
	r.items[rnd], r.items[r.maxRand-1] = r.items[r.maxRand-1], r.items[rnd]
	r.maxRand--
	if r.maxRand == 0 {
		r.maxRand = r.l
	}
	return t
}

// removeElementAt will remove the given element from list
func (r *randomiser) removeElement(elem string) {
	idx := -1
	for i, val := range r.items {
		if val == elem {
			idx = i
			break
		}
	}
	if idx != -1 {
		r.items = append(r.items[:idx], r.items[idx+1:]...)
		r.l--
		if idx < r.maxRand {
			r.maxRand--
		}
	}
}
