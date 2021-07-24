package test1

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func add(x, y int) int {
	return x + y
}
