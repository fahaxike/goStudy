package main

import (
	_ "fmt"

	_ "time"
)

func main() {
  
	stu:=new(Stu)
	// test(stu)

}
fun test(a Le ){

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
	Ho
	Ll
	xixi()
}
type Stu struct {
	Name string
	Age  int
}

func (s *Stu) sleep(a int) {
	panic("not implemented") // TODO: Implement
}

func (s *Stu) dadoudou() int {
	panic("not implemented") // TODO: Implement
}

func (s *Stu) eat() {
	panic("not implemented") // TODO: Implement
}
func (s *Stu) haha() {
	panic("not implemented") // TODO: Implement
}
