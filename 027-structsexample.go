package main

import (
	"fmt"
)

type Uservalues struct{
	ID int `json: "id"`
	CapID string `json: "cap_id"`
	Scheduled bool `json: "scheduled"`
	Status string `json: "status"`
}


func main(){

	data := [][]Uservalues{}

	for i:= 0; i <= 5; i++{
		value := createStructure2(i)
		//if(len(value) > 0){
			data = append(data, value)
		//}
	}

	//var gCtx gin.Context

	fmt.Println(data)
	//fmt.Println(gCtx.JSON(200. data))

}

func createStructure() []Uservalues {
	data := []Uservalues{}

	for i := 0; i < 10; i++ {

		value := Uservalues{
			ID: i,
			CapID: fmt.Sprintf("CAP_%v",i),
			Scheduled : false,
			Status: "Pending",
		}

		data = append(data, value)
	}
	return data
}

func createStructure2(number int) []Uservalues{
	data := []Uservalues{}
	if(number == 5){
		value := Uservalues{
			ID: number,
			CapID: fmt.Sprintf("CAP_%v",number),
			Scheduled : false,
			Status: "Pending",
		}

		data = append(data, value)
	}
	return data
}