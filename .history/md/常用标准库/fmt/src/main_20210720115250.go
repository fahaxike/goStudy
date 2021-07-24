package main

import (
	"fmt"
	"os"
)

func main() {
	// fileObj, err := os.Create("./12.txt")
	fileObj, err := os.OpenFile("d:/1234.txt", os.O_CREATE|os.O_APPEND, 777)
	defer func() {
		fileObj.Close()
		fmt.Println("文件关闭")
	}()
	if err != nil {
		fmt.Printf("111%v", err)
	} else {
		fileObj.WriteString("hahah")
	}

}
