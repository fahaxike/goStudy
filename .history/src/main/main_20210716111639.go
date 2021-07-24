package main

import (
	_ "fmt"                     //引用包但不调用
	hh "goStudy/src/base/test1" //给包起别名
)

func main() {
	hh.Printfmt1()
	hh.Printfmt2()

}
