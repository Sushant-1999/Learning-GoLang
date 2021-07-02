package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Iterating through for loop
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// Iterating with the help of range keyword

	for i, ele := range a {
		// i -> index of slice  (i)
		// ele -> element at that index  (a[i])
		fmt.Printf("Element at %d is %d", i, ele)

	}

	// Another method
	for _, ele := range a {
		fmt.Println(ele)

	}
}
