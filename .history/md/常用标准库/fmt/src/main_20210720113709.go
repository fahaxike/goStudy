package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("D:test.txt", os.O_CREATE|os.O_EXCL|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(fileObj, "xieruleyihang")
	}

}
