package main

import (
	g "computational_geometry/geometry"
	"fmt"
)

func main() {
	a := g.Point{X: 0, Y: 0}
	b := g.Point{X: 2, Y: 2}
	c := g.Point{X: 0, Y: 1}
	d := g.Point{X: 2, Y: 3}

	println(g.Intersect(a, b, c, d))

	polygon := []g.Point{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 0, Y: 4},
	}
	point := g.Point{X: 5, Y: 2}

	distance := g.DistanceToPolygon(polygon, point)
	fmt.Printf("Расстояние от точки до многоугольника: %.2f\n", distance)

	A, B, C := g.Point{X: 0, Y: 0}, g.Point{X: 4, Y: 0}, g.Point{X: 0, Y: 4}

	// Точка внутри
	fmt.Println(g.IsPointInTriangle(A, B, C, g.Point{X: 1, Y: 1}))

	// Точка снаружи
	fmt.Println(g.IsPointInTriangle(A, B, C, g.Point{X: 5, Y: 5}))

	// Точка на границе
	fmt.Println(g.IsPointInTriangle(A, B, C, g.Point{X: 2, Y: 0}))
}
