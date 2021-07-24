package main

import (
	"fmt"

	_ "time"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["hah"] = 1
	m1["hehe"] = 2
	fmt.Printf("type:%T value:%v\n", m1, m1)
	delete(m1, "hehe")
	v, ok := m1["caca"]
	if !ok {
		fmt.Println("不存在caca")
	} else {
		fmt.Printf("caca value:%v\n", v)
	}
	findMap("how do you do")
}

type Homer interface {
	eat()
	sleep(int)
	dadoudou() int
}

type Learner interface {
	eat()
	
}
type Student(){
	Name string
	Age int
}
