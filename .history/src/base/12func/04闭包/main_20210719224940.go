package main

import (
	"fmt"

	_ "time"
)

func main() {

	fn := bibaotest()
	fmt.Println(fn(1))
	fmt.Println(fn(2))
	fmt.Println(fn(3))

}

func bibaotest() func(int) int {
	base := 10
	fmt.Println(base)
	return func(x int) {

		return base + x

	}
}
