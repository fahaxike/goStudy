package main

import (
	"goStudy/src/base/14struct/01var/test1"
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
    var student Student = Student{
		Class: 1,
		Score:31,
		Name:"xiaowang",
		Age:5
        Father.Name:"dawang"
		Father.Age:30
	}
	fmt.Printf("type:%T value:%#v\n", student, student)

}

type Student struct{
	Class: int
	Score:int
    test1.Person
	Father:test1.Person
}
