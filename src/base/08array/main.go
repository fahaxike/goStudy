package main

import (
	"fmt"

	_ "time"
)

func main() {

	var a0 [3]int = [3]int{1, 2, 3}
	fmt.Printf("type:%T value:%v\n", a0, a0)
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
	modify1(a1)
	fmt.Printf("type:%T value:%v\n", a1, a1)
	modify2(&a2)
	fmt.Printf("type:%T value:%v\n", a2, a2)

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
