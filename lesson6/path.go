package main

import (
	"math"
	"fmt"
)

type PointOne struct {
	X, Y float64
}

func (p *PointOne)Distance(q PointOne) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Distance(path []PointOne) float64 {
	source := 0.0
	for i:=0;i<len(path)-1;i++  {
		ok:=path[i].Distance(path[i+1])
		source = source +ok
	}
	return source
}

func main() {
	path := []PointOne{{1,2}, {3, 4}, {5,6}}
	fmt.Println(Distance(path))
}
