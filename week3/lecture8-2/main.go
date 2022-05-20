package main

import (
	"fmt"
	"math"
)

func main() {

	c1 := NewCircle(15.5)
	c2 := NewCircle(13.2)
	c3 := NewCircle(4.1)
	c4 := NewCircle(18.2)
	c5 := NewCircle(8.4)
	s1 := NewSqare(4.9)
	s2 := NewSqare(7.2)
	s3 := NewSqare(5.4)
	s4 := NewSqare(14.8)
	s5 := NewSqare(24.1)
	// fmt.Println(c1)
	// fmt.Println(c1.Area())
	var lotsofshapes Shapes = []Shape{c1, c2, c3, c4, c5, s1, s2, s3, s4, s5}
	fmt.Println(lotsofshapes.LargestArea())
}

type Shape interface {
	Area() float64
}

type Shapes []Shape

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func NewSqare(side float64) Square {
	return Square{side: side}
}
func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}
func (s Square) Area() float64 {
	return s.side * s.side
}
func (r Circle) Area() float64 {
	return r.radius * r.radius * math.Pi
}

func (s Shapes) LargestArea() float64 {
	var Area float64 = 0
	for _, v := range s {
		Area = Area + v.Area()
	}
	return Area
}
