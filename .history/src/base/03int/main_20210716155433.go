package main

import "fmt"

func main() {
	i := 20
	fmt.Printf("i1类型%T\n", i)
	fmt.Printf("i1二进制%b\n", i)
	fmt.Printf("i1十进制%d\n", i)
	fmt.Printf("i1八进制%o\n", i)
	fmt.Printf("i1十六进制%x\n", i)
	i2 := 011
	fmt.Printf("i2类型%T\n", i2)
	fmt.Printf("i2二进制%b\n", i2)
	fmt.Printf("i2十进制%d\n", i2)
	fmt.Printf("i2八进制%o\n", i2)
	fmt.Printf("i2十六进制%x\n", i2)
	i3 := 0x11
	fmt.Printf("i3类型%T\n", i3)
	fmt.Printf("i3二进制%b\n", i3)
	fmt.Printf("i3十进制%d\n", i3)
	fmt.Printf("i3八进制%o\n", i3)
	fmt.Printf("i3十六进制%x\n", i3)

	i4 := int32(i)
	fmt.Printf("i4类型%T\n", i4)
	fmt.Printf("i4二进制%b\n", i4)
	fmt.Printf("i4十进制%d\n", i4)
	fmt.Printf("i4八进制%o\n", i4)
	fmt.Printf("i4十六进制%x\n", i4)

}
