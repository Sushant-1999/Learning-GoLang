package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strconv"
)

// First create an object of scanner
// then use scanner.Scan() and scanner.Text()
// for taking any number of input

func main() {
	scanner := bufio.NewScanner(os.Stdin) // These two lines are compulsory
	scanner.Scan()
	// Input will get stored in scanner
	input := scanner.Text()
	scanner.Scan()
	output := scanner.Text()
	scanner.Scan()
	result := scanner.Text()

	scanner.Scan()
	temp, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Your age is: ", 2021-temp)
	fmt.Println("You typed ", input, output, result)

}
