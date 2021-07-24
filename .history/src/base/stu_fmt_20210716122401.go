package base

import (
	"fmt"
)

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

func Printfmt() {
	name = "xiao名"
	age = 18
	isOk = false
	fmt.Println("hellow world")
}
