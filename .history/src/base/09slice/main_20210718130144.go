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
	fmt.Printf("type:%T value:%v\n", s1, s1)
	fmt.Printf("type:%T value:%v\n", s2, s2)
	fmt.Printf("type:%T value:%v\n", s3, s3)
	fmt.Printf("type:%T value:%v\n", s4, s4)
	fmt.Printf("type:%T value:%v\n", s5, s5)
	fmt.Println()
	a1[1] = 11
	fmt.Printf("type:%T value:%v\n", a1, a1)
	fmt.Printf("type:%T value:%v\n", s1, s1)
	fmt.Printf("type:%T value:%v\n", s2, s2)
	fmt.Printf("type:%T value:%v\n", s3, s3)
	fmt.Printf("type:%T value:%v\n", s4, s4)
	fmt.Printf("type:%T value:%v\n", s5, s5)
	fmt.Println()
	s4[1] = 111
	fmt.Printf("type:%T value:%v\n", a1, a1)
	fmt.Printf("type:%T value:%v\n", s1, s1)
	fmt.Printf("type:%T value:%v\n", s2, s2)
	fmt.Printf("type:%T value:%v\n", s3, s3)
	fmt.Printf("type:%T value:%v\n", s4, s4)
	fmt.Printf("type:%T value:%v\n", s5, s5)
	// fmt.Printf("type:%T value:%v\n", s1, s1)

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
