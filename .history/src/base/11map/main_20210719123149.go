package main

import (
	"fmt"
	"strings"

	_ "time"
)

func main() {

}

func sortMap(a map[string]int) map[string]int {

	fmt.Println(a)
	a[1] = 10
	fmt.Println(a)
	return a
}
func findMap(s string) {
	sl1 := strings.Split(s, " ")
	m1 := make(map[string]int)
	for _, v := range sl1 {
		m1[v] += 1
	}
	for k, v := range m1 {
		fmt.Printf("%s出现了%d次\n", k, v)
	}

}
