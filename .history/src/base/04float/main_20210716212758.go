package main

import (
	"fmt"
	_ "math"
)

func main() {
	// a := math.MaxFloat32
	a1 := 1.2344
	//默认是float64
	fmt.Printf("a1 type %T", a1)
	fmt.Printf("a1  %f,a1 两位 %f", a1, a1)
}
