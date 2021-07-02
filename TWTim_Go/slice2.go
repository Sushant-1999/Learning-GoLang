package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 6, 4, 7, 8, 3, 6}

	for i, element := range a {
		for j, element1 := range a {
			if element == element1 && i > j {
				fmt.Println("Duplicate is:", element)

			}
		}
	}

}
