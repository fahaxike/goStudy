package main

import (
	"fmt"
)

func main() {

	//int2str
	i1 := 55
	s1 := fmt.Sprintf("%d", i1)
	f1 := 5.5
	ss1 := fmt.Sprintf("%.2f", f1)
	fmt.Printf("s1 %T value %s ss1 %T value %s", s1, s1, ss1, ss1)

}
