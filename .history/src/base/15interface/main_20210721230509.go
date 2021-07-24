package main

import (
	_ "fmt"

	_ "time"
)

func main() {

}

type Homer interface {
	eat()
	sleep(int)
	dadoudou() int
}

type Learner interface {
	eat(int)
	Homer
}
type Stu struct {
	Name string
	Age  int
}
