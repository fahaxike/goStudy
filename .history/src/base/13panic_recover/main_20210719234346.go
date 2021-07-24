package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	panic1()
	fmt.Println("2")
}

func panic1() int {
	a := 10
	b := 0
	return a / b

}
