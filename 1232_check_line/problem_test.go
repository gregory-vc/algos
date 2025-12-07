package problem

import "testing"

func TestCheckStraightLine(t *testing.T) {
	res := checkStraightLine([][]int{
		{2, 1},
		{4, 2},
		{6, 3},
	})

	if res != true {
		t.Fatalf("Wrong test")
	}
}

func TestCheckStraightLine2(t *testing.T) {
	res := checkStraightLine2([][]int{
		{2, 1},
		{4, 2},
		{6, 3},
	})

	if res != true {
		t.Fatalf("Wrong test")
	}
}
