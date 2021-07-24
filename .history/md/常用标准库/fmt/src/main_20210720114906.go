package main

import (
	"fmt"
	"os"
)

func main() {
	// fileObj, err := os.Create("./12.txt")
	fileObj, err := os.OpenFile("d:/12.txt", os.O_CREATE|os.O_APPEND, 777)
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("111%v", err)
	} else {
		fmt.Fprintln(fileObj, "xieruleyihang")
	}

}
