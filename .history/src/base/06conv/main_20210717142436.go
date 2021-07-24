package main

import (
	"fmt"
)

func main() {
	// print99()
	// s1 := "c:\\Code\\lesson1\\go.exe"
	// s2 := "hellow北京"
	// fmt.Printf("s1 len:%v,s2 len:%v,s2 count:%v\n", len(s1), len(s2), utf8.RuneCountInString(s2))
	// runes := []rune(s2)
	// fmt.Printf("runes len:%v\n", len(runes))
	// for i, v := range s2 {
	// 	fmt.Printf("s2 i:%v,v:%c,s[i]:%c,\n", i, v, s2[i])
	// }
	// for i, v := range runes {
	// 	fmt.Printf("runes i:%v,v:%c,v:%c\n", i, v, runes[i])
	// }
	// s3 := `asda""sdads\erdf`
	// fmt.Println(s3)
	// s4 := "sd"
	// s4 += "aa"
	// fmt.Println(s4)
	// s4 = "gg" +
	// 	"sdd"
	// fmt.Println(s4)
	// s5 := "54"
	// i5, _ := strconv.Atoi(s5)

	//int2str
	i1 := 55
	s1 := fmt.Sprintf("%d", i1)
	f1 := 5.5
	ss1 := fmt.Sprintf("%2f", f1)
	fmt.Printf("s1 %T value %s sss1 %T value %s", s1, s1, ss1, ss1)

}
