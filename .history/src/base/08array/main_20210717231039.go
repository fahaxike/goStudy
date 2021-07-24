package main

import (
	"fmt"

	_ "time"
)

func main() {

	var a1 = [3]int{1, 2, 3}
	fmt.Printf("type:%T value:%v\n", a1, a1)
	a2 := [3]int{1, 2, 3}
	fmt.Printf("type:%T value:%v\n", a2, a2)
	var a3 = [...]int{1, 2, 3, 4}
	fmt.Printf("type:%T value:%v\n", a3, a3)
	var a4 = [...]int{1: 2, 3: 4}
	fmt.Printf("type:%T value:%v\n", a4, a4)

}
