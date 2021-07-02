package main

import "fmt"

func main() {
	// Arrays in the golang
	// default value initialized is zero(0)
	var arr [10]int
	ptr := []int{} // Initialized as nil
	temp := []string{"Hare", "Krishna", "Rama"}

	// Dealing with multidimensional arrays in golang
	arr2D := [][3]int{{1, 2, 3}, {3, 4, 5}, {7, 8, 9}}
	fmt.Println(arr2D)
	for i := 0; i < len(temp); i++ {
		fmt.Println(temp[i])
	}
	fmt.Println(ptr)
	fmt.Println(arr)

}
