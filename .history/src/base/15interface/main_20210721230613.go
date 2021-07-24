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
	eat()
}
type Stu struct {
	Name string
	Age  int
}
