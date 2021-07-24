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

}

func bibaotest() func(int) int {
	base := 10
	return func(x int) int {

		return base + x

	}
}
