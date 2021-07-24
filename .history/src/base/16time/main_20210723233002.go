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
	unix := now.Unix()
	fmt.Println(unix)
	time.Unix(unix, 0)

	//格式化
	s := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
	fmt.Println(s)
	fmt.Println(now.Format("2006-01-2 3:4:5 "))
	fmt.Println(now.Format("2006-01-2 15:4:5 "))

}
