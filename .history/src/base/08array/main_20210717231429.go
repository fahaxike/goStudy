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
	fmt.Printf("a1等于 a2 is:%v\n", a1 == a2)
	fmt.Printf("a4等于 a3 is:%v\n", a4 == a3)
	var i1 int = 1
	var i2 = 2
	fmt.Printf("type:%T value:%v\n", i1, i1)

}
