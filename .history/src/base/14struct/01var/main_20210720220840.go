package main

import (
	"fmt"

	_ "time"
)

type MyInt int    //自定义类型
type YouInt = int //类型别名
func main() {

	a := MyInt(10)
	b := YouInt(7)
	fmt.Printf("type:%T value:%v\n", a, a)
	fmt.Printf("type:%T value:%v\n", b, b)
	fmt.Printf("type:%T value:%v\n", , )

}
