package test1

import "fmt"

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
type Live interface {
	Shuijiao()

	Study(s string)
}

func add(x, y int) int {
	return x + y
}
func (p *Person) Eat(food string) {
	fmt.Printf("%s正在吃%s", p.Name, food)
}
func (p *Person) Shuijiao() {
	fmt.Printf("%s正在睡觉", p.Name)
}

func (p *Person) Study(s string) {
	fmt.Printf("%s正在学习%s", p.Name, s)
}
