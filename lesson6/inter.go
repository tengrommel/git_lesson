package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p *Point)Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))
}
