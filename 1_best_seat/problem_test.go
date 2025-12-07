package bestseat

import "testing"

func TestMaxDistToClosest(t *testing.T) {
	res := maxDistToClosest([]int{
		1, 0, 0, 0, 1, 0, 1,
	})

	if res != 2 {
		t.Fatalf("Wrong test")
	}
}
