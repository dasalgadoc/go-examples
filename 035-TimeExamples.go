package main

import (
	"fmt"
	"time"
)

func main() {
	myDate := time.Now().Truncate(time.Millisecond).In(time.UTC)
	fmt.Println(myDate)
	timeZone, _ := time.LoadLocation("America/Buenos_Aires")
	myDate2 := time.Now().Truncate(time.Millisecond).In(timeZone)
	fmt.Println(myDate2)
}
