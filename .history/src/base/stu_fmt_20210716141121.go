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

const a = 100
const (
	a1 = 100
	a2
	a3
)
const (
	b1 = 100
	b2 = iota //1
	b3        //2
	b4 = 101  //101
	_         //101
	b6 = iota //3
)

//全局变量可以声明不使用，但是函数内的变量必须声明
func PrintCon() {

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("b1:", b1) //100
	fmt.Println("b2:", b2) //1
	fmt.Println("b3:", b3) //2
	fmt.Println("b4:", b4) //101
	// fmt.Println("b5:", b5)
	fmt.Println("b6:", b6)

}

//全局变量可以声明不使用，但是函数内的变量必须声明
func Printfmt() {
	PrintCon()
	//变量初始化
	// name = "xiao名"
	// age = 18
	// isOk = false
	// fmt.Println(isOk)
	// fmt.Printf("name = %s,age=%v,isOk=%v\r\n", name, age, isOk)
	// //变量初始话
	// var s1 string = "hehe"
	// fmt.Println(s1)
	// //类型推导
	// var s2 = "haha"
	// fmt.Println(s2)
	// //简短声明，只能在函数中使用
	// s3 := "heihei"
	// fmt.Println(s3)

}
