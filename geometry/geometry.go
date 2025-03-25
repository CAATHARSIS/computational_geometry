package geometry

import "math"

type Point struct {
	X, Y float64
}

// Векторное произведение AB на AC
// Если возвращаемое значение больше нуля, то точка C находистя справа от вектора AB
// Если меньше нуля - справа
// Если равно нулю - на прямой
func CrossProduct(A, B, C Point) float64 {
	return (B.X-A.X)*(C.Y-A.Y) - (B.Y-A.Y)*(C.X-A.X)
}

func IsPointOnSegment(A, B, C Point) bool {
	// Проверка на коллиниарность
	if cross := CrossProduct(A, B, C); math.Abs(cross) > 1e-12 {
		return false
	}

	// Проверка на нахождение точки между A и B
	if dot := (B.X-A.X)*(C.X-A.X) + (B.Y-A.Y)*(C.Y-A.Y); dot < 0 {
		return false
	}

	squaredLength := (B.X-A.X)*(B.X-A.X) + (B.Y-A.Y)*(B.Y-A.Y)
	return (C.X-A.X)*(C.X-A.X)+(C.Y-A.Y)*(C.Y-A.Y) <= squaredLength
}

func Intersect(A, B, C, D Point) bool {
	cp1 := CrossProduct(A, B, C)
	cp2 := CrossProduct(A, B, D)
	cp3 := CrossProduct(C, D, A)
	cp4 := CrossProduct(C, D, B)

	if ((cp1 * cp2) < 0) && ((cp3 * cp4) < 0) {
		return true
	}

	if cp1 == 0 && IsPointOnSegment(A, B, C) {
		return true
	}
	if cp2 == 0 && IsPointOnSegment(A, B, D) {
		return true
	}
	if cp3 == 0 && IsPointOnSegment(C, D, A) {
		return true
	}
	if cp4 == 0 && IsPointOnSegment(C, D, B) {
		return true
	}

	return false
}
