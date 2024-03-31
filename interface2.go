package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Shape struct {}

func (this *Shape) Area() float32 {
	panic("Not Impl Error")
	return -1
}

type Square struct {
	side float32
	Shape
}

func (this *Square) Area() float32 {
	return this.side * this.side
}

type Rectangle struct {
	longside float32
	shortside float32
	Shape
}

func (this *Rectangle) Area() float32 {
	return this.longside * this.shortside
}

type Circle struct {
	radius float32
	Shape
}

// func (this *Circle) Area() float32 {
// 	return 0.5 * 3.14159265358979 * this.radius * this.radius
// }

func main() {
	square := Square{side:5}
	rectangle := &Rectangle{shortside:2, longside:3}
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