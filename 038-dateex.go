package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	newTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
	nightlyLte := time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, now.Location())
	fmt.Println(newTime)
	fmt.Println(nightlyLte)
	fmt.Printf("%T", nightlyLte)
}
