package main

import "computational_geometry/geometry"

func main() {
	a := geometry.Point{X: 0, Y: 0}
	b := geometry.Point{X: 2, Y: 2}
	c := geometry.Point{X: 0, Y: 1}
	d := geometry.Point{X: 2, Y: 3}

	println(geometry.Intersect(a, b, c, d))
}
