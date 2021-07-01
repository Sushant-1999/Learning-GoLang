package main

import "fmt"

func main() {
	const temp = 2
	i := 1
	for i <= 5 {
		// Print infinite times
		fmt.Print("Hare Krishna Hare Krishna\n")
		i++
	}

	for i := 1; i <= 100; i++ {
		fmt.Print(i, "\n")
	}

}
