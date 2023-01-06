package main

import (
	"fmt"
	"time"
)

func main() {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)
}
