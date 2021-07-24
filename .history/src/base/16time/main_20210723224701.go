package main

import (
	 "fmt"
	 "time"
)



func main() {
	now:=time.Now()
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hh:=now.Hour()
	ss:=now.Min
	fmt.Println(now)
}
