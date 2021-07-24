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
	fmt.Println(time.Unix(unix, 0))

	//time2str
	s := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
	fmt.Println(s)
	fmt.Println(now.Format("2006-01-2 3:4:5 "))
	fmt.Println(now.Format("2006-01-2 15:4:5 "))
	//str2time
	t1, _ := time.Parse("2006-01-02 15:04:05", "2021-07-23 23:00:00")
	//时间计算
	t1.Add(time.Minute * 5)
	fmt.Println(t1.Add(time.Minute * 5))
	fmt.Println(t1.Add(time.Hour * -10))
	//时间间隔
	fmt.Println(t1.After(time.Now()))
	fmt.Println(t1.Before(time.Now()))

}
