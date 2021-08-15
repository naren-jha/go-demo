package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLen int
}

type rectangle struct {
	b float64
	l float64
}

func (s square) getArea() float64 {
	return float64(s.sideLen * s.sideLen)
}

func (r rectangle) getArea() float64 {
	return r.b * r.l
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{
		sideLen: 10,
	}

	r := rectangle{
		b: 10.1,
		l: 20.2,
	}

	printArea(s) // 100
	printArea(r) // 204.01999999999998
}
