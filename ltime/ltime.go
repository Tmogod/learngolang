package main

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	t0 := time.Now()
	fmt.Printf("%d\n", t0.Unix()) // 当前时间戳
	time.Sleep(50 * time.Millisecond)
	// 时间差
	t1 := time.Now()
	diff := t1.Sub(t0)
	fmt.Printf("%dms\n", diff.Milliseconds())

	// Time + duration = Time
	d := time.Duration(2 * time.Second)
	t2 := t0.Add(d)
	fmt.Printf("%d \n", t2.Unix())

	// 时间的格式化输出
	fmt.Println(t0.Format(DATE))
	s := t0.Format(TIME)
	fmt.Println(t0.Format(TIME))

	t3, _ := time.Parse(TIME, s)
	fmt.Println(t3.Format(TIME))

	// 设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t4, _ := time.ParseInLocation(TIME, s, loc)
	fmt.Println(t4.Format(TIME))
}
