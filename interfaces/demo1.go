package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}
func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şeklin alanı: ", s.Area())

}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)

	c := circle{radius: 10}
	school(c)
}
