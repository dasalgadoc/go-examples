package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum method incorrect! got %d expected %d", total, 10)
	}
}

func TestSumCases(t *testing.T) {
	// Slice de casos de pruebas con los atributos y los resultados esperados
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	// Recorre el slice para probar caso a caso
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum method incorrect! got %d expected %d", total, item.n)
		}
	}
}
