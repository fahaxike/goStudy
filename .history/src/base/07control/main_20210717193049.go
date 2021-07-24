package main

import (
	"fmt"
	"strconv"
	_ "time"

	"golang.org/x/tools/refactor/rename"
)

func main() {

	i := 5
	if i > 5 {
		fmt.Println("i大于5")

	} else if i < 10 {
		fmt.Println("i大于5小于10")
	} else {
		fmt.Println("i大于等于10")
	}
	fmt.Println("---------------------------------------")
	if i := 6; i > 5 {
		fmt.Println("i大于5")

	} else if i < 10 {
		fmt.Println("i大于5小于10")
	} else {
		fmt.Println("i大于等于10")
	}
	fmt.Println("---------------------------------------")
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------------------------")
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------------------------")
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("---------------------------------------")
	fmt.Println(i)

	a := []byte{'a', 'b', 'c', 'd'}
	for k, v := range a {
		fmt.Printf("key:%v value:%c", k, v)
	}
	fmt.Println("---------------------------------------")
	switch i {
	case 5:
		fmt.Println("i=5")
		fallthrough
	case 6:
		fmt.Println("i=6")
		fallthrough
	case 10:
		fmt.Println("i=10")
	case 11:
		fmt.Println("i=11")
	case 12:
		fmt.Println("i=12")
	default:
		fmt.Println("其他")

	}
}
