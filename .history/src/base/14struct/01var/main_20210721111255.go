package main

import (
	_ "encoding/json"
	"fmt"
	"goStudy/src/base/14struct/01var/test1"
	_ "time"
)

type MyInt int    //自定义类型
type YouInt = int //类型别名
var allStu StudentManger

func main() {

	// 	a := MyInt(10)
	// 	b := YouInt(7)
	// 	fmt.Printf("type:%T value:%v\n", a, a)
	// 	fmt.Printf("type:%T value:%v\n", b, b)
	// 	var student Student = Student{
	// 		Class: 1,
	// 		Score: 31,
	// 		// Name:        "xiaowang",
	// 		// Age:         5,
	// 		// Father.Name: "dawang",
	// 		// Father.Age:  30,
	// 		Person: test1.Person{Name: "xiaowang", Age: 10},
	// 		Father: test1.Person{Name: "dawang", Age: 40},
	// 	}
	// 	fmt.Printf("type:%T value:%#v\n", student, student)
	// 	fmt.Printf("type:%T value:%v\n", student, student)
	// 	fmt.Println(student.Name)
	// 	fmt.Println(student.Father.Name)
	// 	var student1 Student

	// 	student.Eat("香蕉")
	// 	student.Father.Eat("苹果")
	// 	j, ok := json.Marshal(student)
	// 	if ok != nil {
	// 		fmt.Println(ok)
	// 	} else {
	// 		fmt.Println(j)
	// 	}
	// 	err := json.Unmarshal(j, &student1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Printf("type:%T value:%#v\n", student1, student1)
	// 	}
	allStu = StudentManger{
		allStudent: make(map[int]Student),
	}
	(&allStu).Add1()
	fmt.Printf("type:%T value:%#v\n", allStu, allStu)
	(&allStu).Add1()
	fmt.Printf("type:%T value:%#v\n", allStu, allStu)

}

type Student struct {
	Class int
	Score int
	ID    int
	test1.Person
	Father test1.Person
}

type StudentManger struct {
	allStudent map[int]Student
	num        int
}

func (s StudentManger) Add() {
	fmt.Println("请输入学生学号 姓名")
	var id int
	var name string
	fmt.Scanf("%d %s  ", &id, &name)
	newStu := Student{
		ID: id,
		Person: test1.Person{
			Name: name,
		},
	}
	s.allStudent[id] = newStu
	s.num++
}
func (s *StudentManger) Add1() {
	fmt.Println("请输入学生学号 姓名")
	var id int
	var name string
	fmt.Scanf("%d %s ", &id, &name)
	newStu := Student{
		ID: id,
		Person: test1.Person{
			Name: name,
		},
	}
	s.allStudent[id] = newStu
	s.num++
}
