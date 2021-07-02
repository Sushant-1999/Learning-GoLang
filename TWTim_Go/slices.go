// Slices in Go
package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5} // Array
	var slc []int
	var slice []int = x[:] // Copied all the elements of an array in the slice
	slc = append(slc, x[3])
	fmt.Println(slc)

	fmt.Println(slice)

}
