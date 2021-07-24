package main

import (
	_ "fmt"               //引用包但不调用
	hh "goStudy/src/base" //给包起别名
)

func main() {
	hh.Printfmt()

}
