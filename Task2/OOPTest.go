package Task2

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func OOPRun() {
	shapes := []Shape{&Rectangle{4, 6}, &Circle{5}}

	rectangle := shapes[0]
	fmt.Printf("Rectangle Area :%f,Perimeter:%f\n", rectangle.Area(), rectangle.Perimeter())

	circle := &Circle{1}
	fmt.Printf("Circle Area :%f,Perimeter:%f\n", circle.Area(), circle.Perimeter())

}
