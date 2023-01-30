package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

// структура Point c 2мя полями x и y типа float64

type Point struct {
	x float64
	y float64
}

// конструктор

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// метод поиска расстояние до конечной точки от текущей. В качестве параметра принимает Point
func (p Point) getDistance(endPoint Point) float64 {
	return math.Sqrt((endPoint.x-p.x)*(endPoint.x-p.x) + (endPoint.y-p.y)*(endPoint.y-p.y))
}

func main() {
	p1 := NewPoint(-6, 2)
	p2 := NewPoint(0, 7)

	fmt.Print(p1.getDistance(p2))
}
