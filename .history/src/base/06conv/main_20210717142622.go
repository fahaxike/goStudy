package main

import (
	"fmt"
)

func main() {

	//int2str
	i1 := 55
	s1 := fmt.Sprintf("%d", i1)
	f1 := 5.5
	ss1 := fmt.Sprintf("%f", f1)
	fmt.Printf("s1 %T value %s ss1 %T value %v", s1, s1, ss1, ss1)

}
