package main

import "fmt"

const a = 100
const (
	a1 = 100
	a2
	a3
	a4 = iota //3
)
const (
	b1 = 100
	b2 = iota //1
	b3        //2
	b4 = 101  //101
	b5        //101
	b6 = iota //5
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)
	fmt.Println("b1:", b1) //100
	fmt.Println("b2:", b2) //1
	fmt.Println("b3:", b3) //2
	fmt.Println("b4:", b4) //101
	fmt.Println("b5:", b5)
	fmt.Println("b6:", b6)
}
