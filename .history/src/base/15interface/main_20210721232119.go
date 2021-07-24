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
	eat(int)
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
}func (s *Stu) sleep(_ int) {
	panic("not implemented") // TODO: Implement
}

func (s *Stu) dadoudou() int {
	panic("not implemented") // TODO: Implement
}

func (s *Stu) eat(_ int) {
	panic("not implemented") // TODO: Implement
}


