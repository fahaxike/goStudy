package main

import (
	_ "fmt"

	_ "time"
)

func main() {

}

type Homer interface {
	sleep(int)
	dadoudou() int
}

type Learner interface {
	eat()
	Homer
}
type Stu struct {
	Name string
	Age  int
}
