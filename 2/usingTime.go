package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time now:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Hour(), t.Minute())
	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("time diff", t1.Sub(t))
	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris time:", londonTime)
}
