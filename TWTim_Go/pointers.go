package main

import "fmt"

func main() {
	x := 7
	y := &x
	// Dereferencing the pointer
	*y = 10

	fmt.Println(x, y)
	fmt.Println(y)

}
