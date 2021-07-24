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
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("收集到错误")
			fmt.Println(err)
		}
	}()
	a := 10
	b := 0
	return a / b

}
