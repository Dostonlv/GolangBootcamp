package main

import (
	"fmt"
	"math"
)

type Implement interface {
	area() float64
	perimetr() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimetr() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) perimetr() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func printImplementInfo(i Implement) {
	fmt.Println("Area:", i.area())
	fmt.Println("Perimeter:", i.perimetr())
}

func main() {

	c := Circle{radius: 2}
	t := Triangle{a: 3, b: 4, c: 5}
	printImplementInfo(c)
	printImplementInfo(t)
}
