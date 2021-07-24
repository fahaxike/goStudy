package main

import (
	"fmt"

	_ "time"
)

func main() {

	a, b := f1(1, 2, "xiao", "da")
	fmt.Println(a, b)

}

func f1(x, y int, s ...string) (r1 int) {
	fmt.Println(s)
	r1 = 0
	return
}

func modify1(a []int) []int {
	fmt.Println(a)
	a[1] = 10
	fmt.Println(a)
	return a
}
func modify2(a *[]int) []int {
	fmt.Println(a)
	(*a)[1] = 15
	fmt.Println(a)
	return *a
}
