package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()

	fmt.Println(now, year, int(month), day, hour, min, sec)
	//时间戳
	s = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
	fmt.Println(s)
}
