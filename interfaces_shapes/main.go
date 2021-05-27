package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{base: 10, height: 10}
	sq := square{sideLength: 4}
	printArea(tr)
	printArea(sq)
}

func (tri triangle) getArea() float64 {
	return tri.base * tri.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
