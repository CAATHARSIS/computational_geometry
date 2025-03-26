package geometry

import "math"

type Point struct {
	X, Y float64
}

// Векторное произведение AB на AC
// Если возвращаемое значение больше нуля, то точка C находистя справа от вектора AB
// Если меньше нуля - справа
// Если равно нулю - на прямой
func crossProduct(A, B, C Point) float64 {
	return (B.X-A.X)*(C.Y-A.Y) - (B.Y-A.Y)*(C.X-A.X)
}

func isPointOnSegment(A, B, C Point) bool {
	// Проверка на коллиниарность
	if cross := crossProduct(A, B, C); math.Abs(cross) > 1e-12 {
		return false
	}

	// Проверка на нахождение точки между A и B (скалярное произведение AC и AB)
	if dot := (B.X-A.X)*(C.X-A.X) + (B.Y-A.Y)*(C.Y-A.Y); dot < 0 {
		return false
	}

	squaredLength := (B.X-A.X)*(B.X-A.X) + (B.Y-A.Y)*(B.Y-A.Y)
	return (C.X-A.X)*(C.X-A.X)+(C.Y-A.Y)*(C.Y-A.Y) <= squaredLength
}

func Intersect(A, B, C, D Point) bool {
	cp1 := crossProduct(A, B, C)
	cp2 := crossProduct(A, B, D)
	cp3 := crossProduct(C, D, A)
	cp4 := crossProduct(C, D, B)

	if ((cp1 * cp2) < 0) && ((cp3 * cp4) < 0) {
		return true
	}

	if cp1 == 0 && isPointOnSegment(A, B, C) {
		return true
	}
	if cp2 == 0 && isPointOnSegment(A, B, D) {
		return true
	}
	if cp3 == 0 && isPointOnSegment(C, D, A) {
		return true
	}
	if cp4 == 0 && isPointOnSegment(C, D, B) {
		return true
	}

	return false
}

func distanceToSegment(P, A, B Point) float64 {
	AB := Point{X: B.X - A.X, Y: B.Y - A.Y}
	AP := Point{X: P.X - A.X, Y: P.Y - A.Y}
	BP := Point{X: P.X - B.X, Y: P.Y - B.Y}

	dotABAP := AB.X*AP.X + AB.Y*AP.Y
	dotABAB := AB.X*AB.X + AB.Y*AB.Y

	// точка P до AB
	if dotABAP <= 0 {
		return math.Sqrt(AP.X*AP.X + AP.Y*AP.Y)
	}
	// точка P после AB
	if dotABAP >= dotABAB {
		return math.Sqrt(BP.X*BP.X + BP.Y*BP.Y)
	}

	return math.Abs(AB.X*AP.Y-AB.Y*AP.X) / math.Sqrt(dotABAB) // AB.X*AP.Y-AB.Y*AP.X - векторное произведение AB x AP
}

func DistanceToPolygon(polygon []Point, P Point) float64 {
	minDist := math.MaxFloat64
	n := len(polygon)

	for i := range n {
		A := polygon[i]
		B := polygon[(i+1)%n]

		dist := distanceToSegment(P, A, B)
		if dist < minDist {
			minDist = dist
		}
	}

	return math.Round(minDist*1000) / 1000
}

func IsPointInTriangle(A, B, C, P Point) bool {
	cp1 := crossProduct(A, B, P)
	cp2 := crossProduct(B, C, P)
	cp3 := crossProduct(C, A, P)

	return (cp1 >= 0 && cp2 >= 0 && cp3 >= 0) || (cp1 <= 0 && cp2 <= 0 && cp3 <= 0)
}
