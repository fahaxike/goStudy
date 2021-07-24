package main

import (
	"fmt"	"math"
)



func main() {
	math.MaxFloat32
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
