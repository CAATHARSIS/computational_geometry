package geometry

import "testing"

func TestIntersect(t *testing.T) {
	var testcases = []struct {
		text  string
		input []Point
		want  bool
	}{
		{"Отрезки пересекаются в одной точке (не на концах)", []Point{
			{X: 1, Y: 1},
			{X: 4, Y: 4},
			{X: 1, Y: 4},
			{X: 4, Y: 1},
		}, true},
		{"Отрезки не пересекаются (параллельны и не лежат на одной линии)", []Point{
			{X: 1, Y: 1},
			{X: 4, Y: 1},
			{X: 1, Y: 2},
			{X: 4, Y: 2},
		}, false},
		{"Отрезки не пересекаются (не параллельны, но не пересекаются)", []Point{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 3},
			{X: 4, Y: 4},
		}, false},
		{"Отрезки пересекаются на концах (общая вершина)", []Point{
			{X: 1, Y: 1},
			{X: 3, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 0},
		}, true},
		{"Отрезки лежат на одной прямой и частично перекрываются", []Point{
			{X: 1, Y: 1},
			{X: 4, Y: 1},
			{X: 2, Y: 1},
			{X: 5, Y: 1},
		}, true},
		{"Отрезки лежат на одной прямой, но не перекрываются", []Point{
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 3, Y: 1},
			{X: 4, Y: 1},
		}, false},
		{"Отрезки совпадают", []Point{
			{X: 1, Y: 1},
			{X: 3, Y: 3},
			{X: 1, Y: 1},
			{X: 3, Y: 3},
		}, true},
		{"Отрезки перпендикулярны, но не пересекаются", []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 4},
			{X: 2, Y: 2},
			{X: 4, Y: 2},
		}, false},
		{"Один отрезок вырожден в точку и лежит на другом отрезке", []Point{
			{X: 1, Y: 1},
			{X: 3, Y: 3},
			{X: 2, Y: 2},
			{X: 2, Y: 2},
		}, true},
		{"Один отрезок вырожден в точку и не лежит на другом отрезке", []Point{
			{X: 1, Y: 1},
			{X: 3, Y: 3},
			{X: 4, Y: 4},
			{X: 4, Y: 4},
		}, false},
		{"Отрезки параллельны и лежат на одной прямой, но касаются только концами", []Point{
			{X: 1, Y: 1},
			{X: 3, Y: 3},
			{X: 3, Y: 3},
			{X: 5, Y: 5},
		}, true},
		{"Отрезки перпендикулярны и пересекаются", []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 4},
			{X: 0, Y: 2},
			{X: 3, Y: 2},
		}, true},
	}

	for _, tt := range testcases {
		t.Run(tt.text, func(t *testing.T) {
			res := Intersect(tt.input[0], tt.input[1], tt.input[2], tt.input[3])

			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}

func TestDistanceToPolygon(t *testing.T) {
	var testcases = []struct {
		text          string
		input_polygon []Point
		input_point   Point
		want          float64
	}{
		{"Простой выпуклый многоугольник (квадрат)", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 3, Y: 1}, 1},
		{"Многоугольник с горизонтальной стороной", []Point{
			{X: 1, Y: 1},
			{X: 4, Y: 1},
			{X: 4, Y: 3},
			{X: 1, Y: 3},
		}, Point{X: 2, Y: 0}, 1},
		{"Многоугольник с вертикальной стороной", []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 4},
			{X: 3, Y: 4},
			{X: 3, Y: 1},
		}, Point{X: 0, Y: 2}, 1},
		{"Точка на продолжении стороны, но вне многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
		}, Point{X: 3, Y: 0}, 1},
		{"Точка ближе к вершине, чем к стороне", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: -1, Y: -1}, 1.414},
		{"Многоугольник с 'вогнутой' частью", []Point{
			{X: 0, Y: 0},
			{X: 3, Y: 0},
			{X: 3, Y: 3},
			{X: 1, Y: 3},
			{X: 1, Y: 1},
			{X: 0, Y: 1},
		}, Point{X: 2, Y: 2}, 1},
		{"Точка на большом удалении от сложного многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
			{X: 0, Y: 1},
		}, Point{X: 10, Y: 10}, 12.728},
		{"Многоугольник с острым углом", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
		}, Point{X: 1, Y: -1}, 1},
		{"Точка близко к 'внутреннему углу' невыпуклого многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 3, Y: 0},
			{X: 3, Y: 3},
			{X: 1, Y: 3},
			{X: 1, Y: 1},
			{X: 0, Y: 1},
		}, Point{X: 0.5, Y: 0.5}, 0.5},
	}

	for _, tt := range testcases {
		t.Run(tt.text, func(t *testing.T) {
			res := DistanceToPolygon(tt.input_polygon, tt.input_point)

			if res != tt.want {
				t.Errorf("got %f, want %f", res, tt.want)
			}
		})
	}
}

func TestIsPointInTriangle(t *testing.T) {
	var testcases = []struct {
		text  string
		input []Point
		want  bool
	}{
		{"Точка внутри треугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 1, Y: 1},
		}, true},
		{"Точка снаружи треугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 3, Y: 1},
		}, false},
		{"Точка на вершине треугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 2, Y: 0},
		}, true},
		{"Точка на ребре треугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 1, Y: 0},
		}, true},
		{"Точка на продолжении ребра, но вне треугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 3, Y: 0},
		}, false},
		{"Точка близко к границе, но снаружи", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 2},
			{X: 1, Y: -0.1},
		}, false},
		{"Вырожденный треугольник (все точки на одной прямой)", []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 1.5, Y: 1.5},
		}, true},
	}

	for _, tt := range testcases {
		t.Run(tt.text, func(t *testing.T) {
			res := IsPointInTriangle(tt.input[0], tt.input[1], tt.input[2], tt.input[3])

			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}

func TestIsPointInPolygon(t *testing.T) {
	var testcaeses = []struct {
		text    string
		polygon []Point
		point   Point
		want    bool
	}{
		{"Простой выпуклый многоугольник (четырёхугольник)", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 1, Y: 1}, true},
		{"Простой невыпуклый многоугольник («звезда»)", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 4, Y: 0},
			{X: 2, Y: 1},
		}, Point{X: 1, Y: 0.5}, true},
		{"Точка снаружи многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 3, Y: 3}, false},
		{"Точка на ребре многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 1, Y: 0}, true},
		{"Точка на вершине многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 2, Y: 0}, true},
		{"Точка на вершине многоугольника", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 2, Y: 0}, true},
		{"Точка близко к границе, но снаружи", []Point{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		}, Point{X: 1, Y: -0.1}, false},
		{"Сложный невыпуклый многоугольник", []Point{
			{X: 0, Y: 0},
			{X: 3, Y: 0},
			{X: 3, Y: 1},
			{X: 2, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 2},
			{X: 3, Y: 3},
			{X: 0, Y: 3},
		}, Point{X: 1, Y: 1.5}, true},
	}

	for _, tt := range testcaeses {
		t.Run(tt.text, func(t *testing.T) {
			res := IsPointInPolygon(tt.polygon, tt.point)

			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}
