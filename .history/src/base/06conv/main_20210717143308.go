package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int2str
	i1 := 55
	s1 := fmt.Sprintf("%d", i1)
	//float2str
	f1 := 5.555
	ss1 := fmt.Sprintf("%.2f", f1)
	fmt.Printf("s1 %T value %s ss1 %T value %s", s1, s1, ss1, ss1)
	s2 := strconv.Itoa(i1)
	ss2 := strconv.FormatFloat(f1, 'f', 1, 64)
	fmt.Printf("s1 %T value %s ss1 %T value %s", s2, s2, ss2, ss2)

}
