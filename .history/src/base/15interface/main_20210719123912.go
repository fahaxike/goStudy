package main

import (
	"fmt"
	"strings"

	_ "time"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["hah"] = 1
	m1["hehe"] = 2
	fmt.Printf("type:%T value:%v\n", m1, m1)
	delete(m1, "hehe")
	v, ok := m1["caca"]
	if !ok {
		fmt.Println("不存在caca")
	} else {
		fmt.Printf("caca value:%v\n", v)
	}
	findMap("how do you do")
}

func sortMap(a map[string]int) {

	fmt.Println(a)
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
