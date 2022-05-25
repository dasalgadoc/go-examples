package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

// Simular un acceso a base de datos, simulación de la función
var GetPersonByDNI = func(dmi string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM ...
	return Person{}, nil
}

// Simular un acceso a base de datos
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil
}
