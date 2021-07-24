package test1

import "fmt"

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
type Live interface {
}

func add(x, y int) int {
	return x + y
}
func (p *Person) Eat(food string) {
	fmt.Printf("%s正在吃%s", p.Name, food)
}
