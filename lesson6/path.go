package main

import (
	"math"
	"fmt"
)

type PointOne struct {
	X, Y float64
}

func (p PointOne)Distance(q PointOne) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []PointOne

func (path Path)Distance() float64 {
	return 0
}

func (path Path)LenPoints() int {
	return len(path)
}

func main() {
	path := Path{{1,2}, {2, 3}, {3, 4}}
	fmt.Println(path.Distance())
}
