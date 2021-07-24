package main

import (
	"bufio"
	_ "bufio"
	_ "errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// fileObj, err := os.Create("./12.txt")
	// fileObj, err := os.OpenFile("d:/123456.txt", os.O_CREATE|os.O_APPEND, 644)
	// defer func() {
	// 	fileObj.Close()
	// 	fmt.Println("文件关闭")
	// }()
	// if err != nil {
	// 	fmt.Printf("111%v", err)
	// } else {
	// 	fileObj.WriteString("hahah")
	// }

	// //写文件2
	// buf := bufio.NewWriter(fileObj)
	// defer buf.Flush()
	// buf.WriteString("hahaha我懂啊so")

	// //readfile1
	// str, _ := os.ReadFile("./main.go")
	// fmt.Println(string(str))
	read3()
}

func read2() {
	//readfile2
	f, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 128)
	var content []byte
	for {
		n, err := f.Read(data)

		if err != nil {
			if err == io.EOF {
				//fmt.Println("读完了")
				break
			}
			fmt.Println(err)
			return
		}
		content = append(content, data[:n]...)
	}

	fmt.Println(string(content))
}

func read3() {
	//readfile2
	f, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	f1 := bufio.NewReader(f)

	data := make([]byte, 128)
	var content []byte
	for {
		line, b, err := f1.ReadLine()

		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println(string(line))
	}

}
