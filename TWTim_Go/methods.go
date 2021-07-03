package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age

}

// For permanent changes we need to pass reference and dereferenct it with pointer as an argument and function passing address of that variable

// Most of the methods need pointer for modifying
func (s *Student) setAge(age int) {
	s.age = age

}

// Getting average grades
func (s Student) getAverage() float32 {
	sum := 0
	for _, v := range s.grades {
		sum = sum + v
	}
	return float32(sum) / float32(len(s.grades))

}

func main() {
	s1 := Student{"Sushant", []int{1, 2, 3, 4, 5}, 22}
	temp := s1.getAge()
	fmt.Println(temp)
	fmt.Println(s1)
	s1.setAge(23)
	fmt.Println(s1)

	avg := s1.getAverage()
	fmt.Println(avg)

}
