package pkg

import "testing"

func Test_randomiser(t *testing.T) {
	r := newRandomiser([]string{"a", "b", "c"})
	a := r.getRandomElement()
	b := r.getRandomElement()
	c := r.getRandomElement()

	if a == b || a == c || b == c {
		t.Error("random generator is not returning unique")
	}

	r.removeElement(c)

	for i := 0; i < 50; i++ {
		if a := r.getRandomElement(); a == c {
			t.Error("removeElementAt did not removed the element")
		}
	}
}
