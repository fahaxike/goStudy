package main

import (
	"fmt"

	_ "time"
)

func main() {
	i1 := []int{1, 2, 3}
	fmt.Println(mapper(i1, double))
	reduce(i1, add, 0)

}
func mapper(i []int, fn func(int) int) []int {
	// result:= make([]int,0,10)
	var result []int
	for _, v := range i {
		fmt.Printf("len:%d,cap:%d\n", len(result), cap(result))
		result = append(result, fn(v))
		fmt.Printf("len:%d,cap:%d\n", len(result), cap(result))
	}
	return result

}
func reduce(i []int, fn func(int, int) int, init int) int {
	fmt.Println(i, init)
	if len(i) <= 0 {
		return init
	}
	return reduce(i[1:], fn, fn(i[0], init))
}
func double(x int) int {
	return 2 * x
}
func add(x, y int) int {
	return x + y
}
