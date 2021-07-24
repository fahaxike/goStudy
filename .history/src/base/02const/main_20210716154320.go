package main

import "fmt"

//单独声明
// var name int
// var age string
// var isOk bool

//一起声明
var (
	name string
	age  int
	isOk bool
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
