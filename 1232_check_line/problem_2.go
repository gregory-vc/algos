package problem

func checkStraightLine2(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	A := Point{
		x: float32(coordinates[0][0]),
		y: float32(coordinates[0][1]),
	}

	B := Point{
		x: float32(coordinates[1][0]),
		y: float32(coordinates[1][1]),
	}

	for _, point := range coordinates[2:] {
		if !checkTriangle(A, B, Point{
			x: float32(point[0]),
			y: float32(point[1]),
		}) {
			return false
		}
	}

	return true
}

func checkTriangle(A, B, C Point) bool {
	return (B.x-A.x)*(C.y-A.y) == (B.y-A.y)*(C.x-A.x)
}
