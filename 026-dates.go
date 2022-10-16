package main

import (
	"fmt"
	"time")

func main(){
	location,_ := time.LoadLocation("America/Sao_Paulo")
	////fmt.Println(date.Now().In(location))
	//hour := time.Now().Truncate(time.Millisecond).In(location)
	//fmt.Println(hour)
//
	//lowerRange := time.Date(2022, 8, 23, 20, 0, 0, 0, location)
	//maxRange := time.Date(2022, 8, 23, 23, 59, 59, 999999999, location)
	//fmt.Println("---------------------")
	//fmt.Printf("%s - %s \n", lowerRange, maxRange)
	//fmt.Println("---------------------")
	//fmt.Println(hour.After(lowerRange))
	//fmt.Println(hour.Before(maxRange))
	//return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).Truncate(time.Millisecond)
	owner := "662425803"
	accum := "32499120-57f4-43ae-a3d8-0eccbe600951"
	myDate := time.Date(2022, 10, 5, 0, 0, 0, 0, location).Truncate(time.Millisecond)
	myString := fmt.Sprintf("%s_%s_%v",owner,accum,myDate,)
	fmt.Println(myString)
}