package main

import (
	 "fmt"

	 "time"
)



func main() {
	now:=time.Now()
	year:=now.year
	month:=now.month
	day:=now.day
	hh:=now.hour
	ss:=now.min
	fmt.Println(now)
}
