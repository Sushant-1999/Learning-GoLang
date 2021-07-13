package main

import (
	"fmt"
	"math"
)

type geometry interface {
	// collection of common methods signatures
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Now implementing methods on these structs(shapes)

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

// Now perimeter methods on these shapes

func (r rect) perim() float64 {
	return 2 * r.width * r.height

}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius

}

// Now function for accessing the methods defined in
// the interface geometry
// this is possible when we assign a variable to that interface

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	r := rect{10, 11}
	c := circle{10}
	measure(r)
	measure(c)

}
