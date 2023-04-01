package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Persona struct {
	Nombre  string
	Medidas Medidas
}

type Medidas struct {
	Altura float64
	Peso   float64
}

func main() {
	m := Medidas{
		Altura: 1.70,
		Peso:   70.0,
	}

	p := Persona{
		Nombre:  "Diego",
		Medidas: m,
	}

	path := "Medidas"
	fieldValue := getFieldValue(p, path)
	//fmt.Println(fieldValue.Float())
	fmt.Println(fieldValue)
	if !fieldValue.IsValid() {
		fmt.Println("No valido")
	}
}

func getFieldValue(data interface{}, path string) reflect.Value {
	fields := strings.Split(path, ".")
	value := reflect.ValueOf(data)
	for _, field := range fields {
		value = value.FieldByName(field)
	}
	return value
}
