package main

import (
	"fmt"
	"os"
)

func main() {
	// fileObj, err := os.Create("./12.txt")
	//读文件

	//写文件1
	fileObj, err := os.OpenFile("d:/123456.txt", os.O_CREATE|os.O_APPEND, 644)
	defer func() {
		fileObj.Close()
		fmt.Println("文件关闭")
	}()
	if err != nil {
		fmt.Printf("111%v", err)
	} else {
		fileObj.WriteString("hahah")
	}
	//写文件2 
	bufio.

}
