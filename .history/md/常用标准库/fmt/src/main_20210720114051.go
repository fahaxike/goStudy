package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./12.txt", os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(fileObj, "xieruleyihang")
	}

}
