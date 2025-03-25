package main

import (
	"computational_geometry/geometry"
	"fmt"
)

func main() {
	a := geometry.Point{X: 0, Y: 0}
	b := geometry.Point{X: 2, Y: 2}
	c := geometry.Point{X: 0, Y: 1}
	d := geometry.Point{X: 2, Y: 3}

	println(geometry.Intersect(a, b, c, d))

	polygon := []geometry.Point{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 0, Y: 4},
	}
	point := geometry.Point{X: 5, Y: 2}

	distance := geometry.DistanceToPolygon(polygon, point)
	fmt.Printf("Расстояние от точки до многоугольника: %.2f\n", distance)
}
