package problem

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	k, b, isVertical := solve(Point{
		x: float32(coordinates[0][0]),
		y: float32(coordinates[0][1]),
	}, Point{
		x: float32(coordinates[1][0]),
		y: float32(coordinates[1][1]),
	})

	for _, point := range coordinates[2:] {
		if !check(coordinates[0][0], k, b, Point{
			x: float32(point[0]),
			y: float32(point[1]),
		}, isVertical) {
			return false
		}
	}

	return true
}

func solve(point1, point2 Point) (float32, float32, bool) {
	if point1.x == point2.x {
		return 0, 0, true
	}

	k := (point2.y - point1.y) / (point2.x - point1.x)
	b := point1.y - k*point1.x

	return k, b, false
}

func check(x1 int, k, b float32, point Point, isVertical bool) bool {
	if isVertical {
		return float32(x1) == point.x
	}
	return point.y == k*point.x+b
}

type Point struct {
	x float32
	y float32
}
