package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 1}
	q := Point{4, 5}
	fmt.Println(Distance(&q, &p))
	fmt.Println((&q).Distance(&p))
}

type Point struct{ X, Y float64 }

func Distance(p, q *Point) float64 {
	return math.Hypot((*p).X-(*q).X, (*p).Y-(*q).Y)
}

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot((*p).X-(*q).X, (*p).Y-(*q).Y)
}
