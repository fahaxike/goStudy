package main

import (
	_ "fmt"

	_ "time"
)

func main() {

}

type Ho interface {
	sleep(int)
	dadoudou() int
	eat()
}
type Ll interface {
	eat()
	haha()
}
type Le interface {
}
type Stu struct {
	Name string
	Age  int
}
