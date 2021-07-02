package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apples": 5,
		"banana": 4,
		"orange": 3,
	}
	// Initialized a map
	// Order is not sorted in map that's it
	// Don't care of order of data then here comes map

	fmt.Println(mp)

	/*
		mp := make(map[string]int)
	*/

	// deleting entries from maps

	delete(mp, "apples")
	fmt.Println(mp)

	// Now checking if the key exists inside the map

}
