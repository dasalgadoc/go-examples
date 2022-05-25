package main

import (
	"fmt"
)

func simpleProcedure() {
	fmt.Println("Soy un procedimiento simple")
}

func procedureWithParameters(a int, b int, c string) {
	fmt.Printf("Los parámetros son %d %d %s \n", a, b, c)
}

func procedureGoodPractice(a, b int, c string) {
	fmt.Printf("Implemento buena práctica en los parámetros %d %d %s \n", a, b, c)
}

func functionWithReturn() int {
	return 100
}

func functionWithTwoReturns() (a, b int) {
	return 200, 300
}

func functionTwoReturnsName() (mesa, votos int, presidente string) {
	mesa = 8
	votos = 234
	presidente = "Ramiro Ramirez"
	return
}

func sumWithSlice(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	simpleProcedure()
	procedureWithParameters(10, 20, "Diego")
	procedureWithParameters(30, 40, "Daniela")

	var100 := functionWithReturn()
	fmt.Println("La siguiente variable es el retorno de una función: ", var100)

	var200, var300 := functionWithTwoReturns()
	fmt.Println("Las siguientes variables son el retorno de una función: ", var200, var300)

	var200, _ = functionWithTwoReturns()
	fmt.Println("Se ignoró el segundo parámetro")

	_, var300 = functionWithTwoReturns()
	fmt.Println("Se ignoró el primer parámetro")

	fmt.Println(functionTwoReturnsName())

	fmt.Println("Suma de dos valores: ", sumWithSlice(1, 2))
	fmt.Println("Suma de tres valores: ", sumWithSlice(1, 2, 3))
	fmt.Println("Suma de cinco valores: ", sumWithSlice(1, 2, 3, 4, 5))

	// Función anónima

	func() {
		fmt.Println("Hola, soy una función anónima")
	}()

	func(text string) {
		fmt.Printf("Hola, soy una función anónima con este %s parámetro\n", text)
	}("parametro")

	myAnonFunc := func() {
		fmt.Printf("Anon\t\n")
	}

	myAnonFunc()
	myAnonFunc()
	myAnonFunc()

}
