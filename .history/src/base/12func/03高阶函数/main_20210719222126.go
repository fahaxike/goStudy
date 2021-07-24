package main

import (
	"fmt"

	_ "time"
)

func main() {
	i1 := []int{1, 2, 3}
	fmt.Println(mapper(i1, double))

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
func double(x int) int {
	return 2 * x
}
