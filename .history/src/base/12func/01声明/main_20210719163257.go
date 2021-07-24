package main

import (
	"fmt"

	_ "time"
)

var i = 10 //全局变量

func main() {

	a, b := f1(1, 2, "xiao", "da")
	fmt.Println(a)
	fmt.Println(b)
	i := "ss" //局部变量，局部变量优先级大于全局变量
	fmt.Println(i)
}

//相同类型边变量可以只写最后一个，可变参数只能有一个，并且要放在最后
//返回值变量可以声明，return后面可以不写变量
func f1(x, y int, s ...string) (r1 int, r2 string) {
	fmt.Println(s)
	// r1 = 0
	// r2 = "hah"
	fmt.Println(i)
	return
}

func modify1(a []int) []int {
	fmt.Println(a)
	a[1] = 10
	fmt.Println(a)
	return a
}
func modify2(a *[]int) []int {
	fmt.Println(a)
	(*a)[1] = 15
	fmt.Println(a)
	return *a
}
