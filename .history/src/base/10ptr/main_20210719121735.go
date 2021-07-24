package main

import (
	"fmt"
)

func main() {
	//指针地址 &a，指针类型 *int，指针取值 *a
	a := 10
	b := &a
	fmt.Printf("type:%T value:%v address:%v,ptr:%v\n", a, a, &a, *b)
	fmt.Printf("type:%T value:%v address:%v\n", b, b, &b)

	//对于引用类型需要初始化 new用于指针初始化，返回对应类型的指针,make 用于 切片，map,channel初始化 返回的是引用类型本身
	var a1 *int
	a1 = new(int)
	*a1 = 100
	fmt.Printf("type:%T value:%v address:%v\n", a1, a1, &a1)
	var s1 []int = make([]int, 5)
	var m1 map[string]int = make(map[string]int, 10)
	var c1 = make(chan int, 10)
	fmt.Printf("type:%T value:%v\n", s1, s1)
	fmt.Printf("type:%T value:%v \n", m1, m1)
	fmt.Printf("type:%T value:%v\n", c1, c1)
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
