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
	//变量初始化
	name = "xiao名"
	age = 18
	isOk = false
	fmt.Println(isOk)
	fmt.Printf("name = %s,age=%v,isOk=%v\r\n", name, age, isOk)
	//变量初始话
	var s1 string = "hehe"
	fmt.Println(s1)
	//类型推导
	var s2 = "haha"
	fmt.Println(s2)
	//简短声明，只能在函数中使用
	s3 := "heihei"
	fmt.Println(s3)
}
