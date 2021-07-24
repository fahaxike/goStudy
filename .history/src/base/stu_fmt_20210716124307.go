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

//全局变量可以声明不使用，但是函数内的变量必须声明
func Printfmt() {
	//变量初始化
	name = "xiao名"
	age = 18
	isOk = false
	fmt.Println(isOk)
	fmt.Printf("name = %s,age=%v,isOk=%v", name, age, isOk)
	//变量初始话
	var s1 string = "hehe"
	fmt.Println(s1)
	//类型推导
	var s2 = "haha"
	fmt.Println(s2)

	s3 := "heihei"
	fmt.Println(s3)

}
