package main

import "fmt"

type Shap interface {
	area() float32
}

type Rect struct {
	width  float32
	height float32
}

type Circle struct {
	// 半径
	radius float32
}

func (this Rect) area() float32 {
	return this.width * this.height
}

func (this Circle) area() float32 {
	return 3.14 * this.radius * this.radius
}

func main() {
	rect := Rect{width: 10, height: 10}
	fmt.Println(rect.area())

	circle := Circle{radius: 100}
	fmt.Println(circle.area())
}
