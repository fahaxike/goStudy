package main

import (
	"fmt"

	_ "time"
)

func main() {

	var a1 = [2]int{0, 1}
	// s1 := a1[:2] //
	// s2 := a1[1:2]
	// s3 := a1[1:]
	// s4 := a1[:]
	// s5 := s3[:2]
	// fmt.Printf("type:%T value:%v\n", a1, a1)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	// fmt.Printf("type:%T value:%v\n", s2, s2)
	// fmt.Printf("type:%T value:%v\n", s3, s3)
	// fmt.Printf("type:%T value:%v\n", s4, s4)
	// fmt.Printf("type:%T value:%v\n", s5, s5)
	// fmt.Println()
	// a1[1] = 11
	// fmt.Printf("type:%T value:%v\n", a1, a1)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	// fmt.Printf("type:%T value:%v\n", s2, s2)
	// fmt.Printf("type:%T value:%v\n", s3, s3)
	// fmt.Printf("type:%T value:%v\n", s4, s4)
	// fmt.Printf("type:%T value:%v\n", s5, s5)
	// fmt.Println()
	// s4[1] = 111
	// fmt.Printf("type:%T value:%v\n", a1, a1)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	// fmt.Printf("type:%T value:%v\n", s2, s2)
	// fmt.Printf("type:%T value:%v\n", s3, s3)
	// fmt.Printf("type:%T value:%v\n", s4, s4)
	// fmt.Printf("type:%T value:%v\n", s5, s5)
	// modify1(s4)
	// fmt.Printf("type:%T value:%v\n", a1, a1)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	// fmt.Printf("type:%T value:%v\n", s2, s2)
	// fmt.Printf("type:%T value:%v\n", s3, s3)
	// fmt.Printf("type:%T value:%v\n", s4, s4)
	// fmt.Printf("type:%T value:%v\n", s5, s5)
	// fmt.Println()
	// modify2(&s4)
	// fmt.Printf("type:%T value:%v\n", a1, a1)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	// fmt.Printf("type:%T value:%v\n", s2, s2)
	// fmt.Printf("type:%T value:%v\n", s3, s3)
	// fmt.Printf("type:%T value:%v\n", s4, s4)
	// fmt.Printf("type:%T value:%v\n", s5, s5)
	// fmt.Printf("type:%T value:%v\n", s1, s1)
	//当切片扩容之后，小于等于数组本身，切片会更换地址，原本的数组或者切片数值改变之后，扩容的切片元素内容不再改变
	a2 := make([]int, 2)
	fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	a2 = a1[:2]
	fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	a2 = append(a2, 5)
	fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	a1[0] = 10
	fmt.Printf("type:%T value:%v,ptr:%p\n", a1, a1, &a1)
	fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	a2[0] = 11
	fmt.Printf("type:%T value:%v,ptr:%p\n", a1, a1, &a1)
	fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)

	// //当切片扩容之后，大于数组本身，切片会更换地址，原本的数组或者切片数值改变之后，扩容的切片元素内容不再改变
	// a2 := make([]int, 3, 4)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	// a2 = a1[:]
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	// a2 = append(a2, 5)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)
	// a1[0] = 10
	// fmt.Printf("type:%T value:%v,ptr:%p\n", a1, a1, &a1)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, len(a2), cap(a2), a2)

	// a3 := append(a2, 5)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, a2, a2, a2)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a3, a3, a3, a3, a3)
	// a2[0] = 11
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a1, a1, a1, a1, &a1)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a2, a2, a2, a2, a2)
	// fmt.Printf("type:%T value:%v,len:=%v,cap:%v,ptr:%p\n", a3, a3, a3, a3, a3)
	// a2 = append(a2, 5)
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
