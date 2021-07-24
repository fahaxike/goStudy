package main

import (
	"fmt"
	"goStudy/src/base/14struct/01var/test1"
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
		Score: 31,
		// Name:        "xiaowang",
		// Age:         5,
		// Father.Name: "dawang",
		// Father.Age:  30,
		Person: test1.Person{Name: "xiaowang", Age: 10},
		Father: test1.Person{Name: "dawang", Age: 40},
	}
	fmt.Printf("type:%T value:%#v\n", student, student)
	fmt.Printf("type:%T value:%v\n", student, student)
	fmt.Println(student.Name)
	fmt.Println(student.Father.Name)
	var student1 Student = Student{
		Class: 0,
		Score: 0,
		Person: test1.Person{
			Name: "",
			Age:  0,
		},
		Father: test1.Person{
			Name: "",
			Age:  0,
		},
	}

}

type Student struct {
	Class int
	Score int
	test1.Person
	Father test1.Person
}
