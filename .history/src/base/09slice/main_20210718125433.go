package main

import (
	"fmt"

	_ "time"
)

func main() {

	var a1 = [5]int{0, 1, 2, 3, 4}
	s1 := a1[:2]
	s2 := a1[1:2]
	s3 := a1[1:]
	s4 := a1[:]
	s5 := s3[:2]
	fmt.Printf("type:%T value:%v\n", a1, a1)
	fmt.Printf("''\n")
}

func modify1(a [3]int) [3]int {
	fmt.Println(a)
	a[0] = 10
	fmt.Println(a)
	return a
}
func modify2(a *[3]int) [3]int {
	fmt.Println(a)
	a[0] = 10
	fmt.Println(a)
	return *a
}
