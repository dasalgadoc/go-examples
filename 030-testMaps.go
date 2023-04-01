package main

import (
	"fmt"
)

type Project struct {
	Id    int64  `json:"project_id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func main() {

	project := Project{
		Id:    123,
		Title: "Mi titulo",
		Name:  "Mi name",
	}

	fmt.Printf("%+v\n", project)

	keys := map[string]string{
		"LUNES":     "1",
		"MARTES":    "2",
		"MIERCOLES": "3",
		"JUEVES":    "4",
		"VIERNES":   "5",
	}

	dayOfWeek := "MARTES"
	incorrectDayOfWeek := ""

	filterOkDay, ok := keys[dayOfWeek]

	if ok {
		fmt.Println(filterOkDay)
	} else {
		fmt.Println("NO LO ENCONTRÉ")
	}

	filterOkDay, ok = keys[incorrectDayOfWeek]

	if ok {
		fmt.Println(filterOkDay)
	} else {
		fmt.Println("NO LO ENCONTRÉ")
	}

}
