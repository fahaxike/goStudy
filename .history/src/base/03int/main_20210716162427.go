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
	//字面量语法
	i5 := 0b10011
	fmt.Printf("i5类型%T\n", i5)
	fmt.Printf("i5二进制%b\n", i5)
	fmt.Printf("i5十进制%d\n", i5)
	fmt.Printf("i5八进制%o\n", i5)
	fmt.Printf("i5十六进制%x\n", i5)
	//字面量语法
	i6 := 0o11
	fmt.Printf("i6类型%T\n", i6)
	fmt.Printf("i6二进制%b\n", i6)
	fmt.Printf("i6十进制%d\n", i6)
	fmt.Printf("i6八进制%o\n", i6)
	fmt.Printf("i6十六进制%x\n", i6)
	//字面量语法
	i7 := 0x11
	fmt.Printf("i7类型%T\n", i7)
	fmt.Printf("i7二进制%b\n", i7)
	fmt.Printf("i7十进制%d\n", i7)
	fmt.Printf("i7八进制%o\n", i7)
	fmt.Printf("i7十六进制%x\n", i7)
}
