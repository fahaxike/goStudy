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

type Le interface {
	Homer
}
type Stu struct {
	Name string
	Age  int
}
