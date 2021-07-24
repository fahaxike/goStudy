package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "c:\\Code\\lesson1\\go.exe"
	s2 := "hellow北京"
	fmt.Printf("s1 len:%v,s2 len:%v,s2 count:%v\n", len(s1), len(s2), utf8.RuneCountInString(s2))
	runes := []rune(s2)
	fmt.Printf("runes len:%v\n", len(runes))
	for i, v := range s2 {
		fmt.Printf("s2 i:%v,v:%c,s[i]:%c,\n", i, v, s2[i])
	}
	for i, v := range runes {
		fmt.Printf("runes i:%v,v:%c,v:%c\n", i, v, runes[i])
	}

}
