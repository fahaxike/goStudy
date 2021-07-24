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
	f1 := 5.5555
	ss1 := fmt.Sprintf("%.3f", f1)
	fmt.Printf("s1 %T value %s ;ss1 %T value %s\n", s1, s1, ss1, ss1)
	s2 := strconv.Itoa(i1)
	ss2 := strconv.FormatFloat(f1, 'f', 3, 64)
	fmt.Printf("s2 %T value %s; ss2 %T value %s\n", s2, s2, ss2, ss2)
	//bool2str
	b1 := true
	sb1 := fmt.Sprintf("%t", b1)
	ssb1 := strconv.FormatBool(b1)
	fmt.Printf("sb1 %T value %s; ssb1 %T value %s\n", sb1, sb1, ssb1, ssb1)

}
