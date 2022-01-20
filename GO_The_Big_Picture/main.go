package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hare.txt")
	if err != nil {
		log.Fatal("Opening File failed")
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		} else {
			fmt.Println(s)
		}
	}
}
