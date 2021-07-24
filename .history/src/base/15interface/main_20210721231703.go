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
	eat()
}

type Le interface {
	Homer
	eat()
}
type Stu struct {
	Name string
	Age  int
}
