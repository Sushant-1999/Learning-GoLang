package main

import (
	"fmt"
)

// Separating out the Business Logic
// yes we can return mulitple value from functions
// No need to specify return this and this if specified at time of return definition

func add(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return

}

func main() {
	num1 := 10
	num2 := 20

	result1, result2 := add(num1, num2)
	fmt.Print(result1, result2)

}
