package main

import "fmt"

type MyStruct struct {
	Attr string
}

type Flowable interface{}

var ImplementedFlows = map[string]Flowable{
	"base": &MyStruct{},
}

var ImplementedFlows2 = map[string]func() Flowable{
	"base": GetStruct,
}

func GetStruct() Flowable {
	return &MyStruct{}
}

func main() {
	//flow, _ := ImplementedFlows["base"]
	//
	//flow2, _ := ImplementedFlows["base"]
	//
	//fmt.Println(flow)
	//fmt.Println(flow2)
	//
	//fmt.Printf("%p\n",flow)
	//fmt.Printf("%p\n",flow2)
	//
	//flow3 := ImplementedFlows2["base"]()
	//flow4 := ImplementedFlows2["base"]()
	//
	//fmt.Println(flow3)
	//fmt.Println(flow4)
	//
	//fmt.Printf("%p\n",flow3)
	//fmt.Printf("%p\n",flow4)

	var flows []Flowable

	for i := 0; i < 10; i++ {
		flows = append(flows, ImplementedFlows2["base"]())
	}

	for _, flow := range flows {
		fmt.Println(flow)
		fmt.Printf("%p\n", flow)
	}
}
