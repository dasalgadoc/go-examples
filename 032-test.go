package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

const (
	pix = "PIX"
)

var (
	flows = []string{pix}
)

type ClientResponse map[string]interface{}

type DynamicParam struct {
	param string
	value string
}

type Persona struct {
	Nombre string
	Edad   int
}

type SimpleStruct struct {
	aString  string
	aInt     int
	aBoolean bool
}

type NestedStruct struct {
	title  string
	nested SimpleStruct
}

type Person struct {
	Name   string
	Age    int
	Gender string
	Single bool
}

func main() {
	testMarshall()
}

func testMarshall() {
	var clientResponse ClientResponse

	pi := 3.14
	s := `{"valor_pi": %.2f}`
	b := []byte(fmt.Sprintf(s, pi))

	err := json.Unmarshal(b, &clientResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(clientResponse)
}

func testExceptionAtoi() {
	number := "123A"
	converted, err := strconv.Atoi(number)

	fmt.Println(converted, err.Error())
	ss := fmt.Sprintf(`0 strconv.Atoi: parsing "%s": invalid syntax`, number)
	fmt.Println(ss)
	if err.Error() == ss {
		fmt.Println("Hola!")
	}
}

func testFuncs() {
	ubay := Person{
		Name:   "John",
		Gender: "Female",
		Age:    17,
		Single: false,
	}
	values := reflect.ValueOf(ubay)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}
}

func testMapsWithFunctions() {
	var name string
	name = "Nico"
	switch name {
	case "Diego":
		sayHelloDiego()
	case "Nico":
		sayHelloNico()
	}

	m := map[string]func(){
		"Diego": sayHelloDiego,
		"Nico":  sayHelloNico,
	}
	m[name]()
}

func sayHelloDiego() {
	fmt.Println("Hola Diego")
}

func sayHelloNico() {
	fmt.Println("Hola Nico")
}

func testDynamic() {
	site := DynamicParam{param: "siteid", value: "mlb"}
	resp := NewClientResponse(site)

	persona := Persona{Nombre: "Diego", Edad: 31}
	resp2 := NewClientResponse(persona)

	fmt.Println(resp)
	fmt.Println(resp2)

}

func NewClientResponse(s interface{}) ClientResponse {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	m := ClientResponse{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name)
		m[f.Name] = v.Field(i)
	}
	return m
}

func testLoop() {
	for _, v := range flows {
		if v == "pix" {
			fmt.Println("Hola")
		}
	}
}

func testDynamicStructs() {
	dynamicParams := []DynamicParam{}

	site := DynamicParam{param: "siteid", value: "mlb"}
	owner := DynamicParam{param: "ownerid", value: "123"}

	dynamicParams = append(dynamicParams, site, owner)
	fmt.Println(dynamicParams)
}
