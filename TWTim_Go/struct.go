package main

import "fmt"

// Making a structure
// Custom type that we can use
type Point struct {
	// Fields of struct
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{3, 4}
	fmt.Println("P2 before", p2)
	changeX(&p2)
	fmt.Println("P2 after", p2)

	fmt.Println(p1.x)
	fmt.Println(p2.y)

}
