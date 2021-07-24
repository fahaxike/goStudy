package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// //int2str
	// i1 := 55
	// s1 := fmt.Sprintf("%d", i1)
	// //float2str
	// f1 := 5.5555
	// ss1 := fmt.Sprintf("%.3f", f1)
	// fmt.Printf("s1 %T value %s ;ss1 %T value %s\n", s1, s1, ss1, ss1)
	// s2 := strconv.Itoa(i1)
	// ss2 := strconv.FormatFloat(f1, 'f', 3, 64)
	// fmt.Printf("s2 %T value %s; ss2 %T value %s\n", s2, s2, ss2, ss2)
	// //bool2str
	// b1 := true
	// sb1 := fmt.Sprintf("%t", b1)
	// ssb1 := strconv.FormatBool(b1)
	// fmt.Printf("sb1 %T value %s; ssb1 %T value %s\n", sb1, sb1, ssb1, ssb1)
	//str2int
	s1 := "55"
	s2 := "55.5556"
	s3 := "true"
	i1, _ := strconv.Atoi(s1)
	ii1, _ := strconv.ParseInt(s1, 16, 64)
	f2, _ := strconv.ParseFloat(s2, 64)
	b3, _ := strconv.ParseBool(s3)
	fmt.Printf("i1 %T value %v; \n", i1, i1)
	fmt.Printf("ii1 %T value %v; \n", ii1, ii1)
	fmt.Printf("f2 %T value %v; \n", f2, f2)
	fmt.Printf("b3 %T value %v; \n", b3, b3)
	var name string
	var age int
	var score int
	var ggg int
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 成绩")
	fmt.Scanf("%d %d ", &age, &score)
	fmt.Printf("姓名：%s 年龄：%d 成绩：%d\n", name, age, score)
	fmt.Println("===")
	fmt.Scanln(&ggg)
	fmt.Printf("姓名：%s 年龄：%d 成绩：%d", name, age, ggg)
	time.Sleep(time.Second * 5)
	//str2float
	//str2bool

}
