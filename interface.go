package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (this *Square) Area() float32 {
	return this.side * this.side
}

type Rectangle struct {
	longside float32
	shortside float32
}

func (this *Rectangle) Area() float32 {
	return this.longside * this.shortside
}

type Circle struct {
	radius float32
}

func (this *Circle) Area() float32 {
	return 0.5 * 3.14159265358979 * this.radius * this.radius
}

func main() {
	square := Square{5}
	rectangle := &Rectangle{2, 3}
	circle := new(Circle)
	circle.radius = 3

	var areaIntf Shaper

	areaIntf = &square
	fmt.Println(areaIntf.Area())
	areaIntf = rectangle
	fmt.Println(areaIntf.Area())
	areaIntf = circle
	fmt.Println(areaIntf.Area())
}