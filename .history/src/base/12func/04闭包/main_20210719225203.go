package main

import (
	"fmt"

	_ "time"
)

func main() {

	fn := bibaotest()
	fmt.Println(fn(1))
	fmt.Println(fn(1))
	fmt.Println(fn(1))
	fn1 := bibaotest1()
	fmt.Println(fn(1))
	fmt.Println(fn(1))
	fmt.Println(fn(1))

}

func bibaotest() func(int) int {
	base := 10
	return func(x int) int {
		base = base + x
		return base

	}
}
func bibaotest1() func(int) int {
	base := 10
	return func(x int) int {
		return base + x

	}
}
