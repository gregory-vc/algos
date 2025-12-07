package bestseat

func maxDistToClosest(seats []int) int {
	max := 0
	current := 0
	firstOne := -1
	lastOne := -1
	for k, v := range seats {
		if v == 0 {
			current++
		} else {
			if firstOne < 0 {
				firstOne = k
			}
			lastOne = k
			max = maximum(current, max)
			current = 0
		}
	}

	middleOne := (max + 1) / 2
	lastOne = len(seats) - 1 - lastOne

	max = maximum(middleOne, lastOne)
	return maximum(max, firstOne)
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
