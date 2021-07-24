package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("1")
	a, err := panic1(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
	fmt.Println("2")

}

func panic1(a, b int) (int, error) {

	if b == 0 {
		err := errors.New("除数不能为能0")
		return 0, err
	} else {
		return a / b, nil
	}

}
