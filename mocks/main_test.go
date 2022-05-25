package main

import (
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEMployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{Name: "Andy", Age: 35, DNI: "1"}, nil
				}
			},
			expectedEMployee: FullTimeEmployee{
				Person:   Person{Name: "Andy", Age: 35, DNI: "1"},
				Employee: Employee{Id: 1, Position: "CEO"},
			},
		},
	}

	// Guardar las funciones originales
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI

	for _, item := range table {
		item.mockFunc()
		ft, err := GetFullTimeEmployeeById(item.id, item.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != item.expectedEMployee.Age {
			t.Errorf("Error in Age, got %d expected %d", ft.Age, item.expectedEMployee.Age)
		}

	}

	// Restore the original functions
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}
