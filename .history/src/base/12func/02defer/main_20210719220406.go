package main

import (
	"fmt"

	_ "time"
)

func main() {
	a := 1
	b := 2
	defer cal(1, a, cal(2, a, b))
	a = 0
	defer cal(3, a, cal(4, a, b))
	b = 1
	//2,1,2,3
	//4,0,2,0
	//1,1,3,4
	//3,0,2,2
}

func cal(index, x, y int) int {
	fmt.Println(index, x, y, x+y)
	return x + y
}
