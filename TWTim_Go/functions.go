package main

import "fmt"

func add(x int, y int) (z1, z2 int) {
	defer fmt.Println("Addition Done")
	z1 = x + y
	z2 = x - y
	return

}

func main() {
	ans1, ans2 := add(14, 7)
	fmt.Println(ans1, ans2)

}
