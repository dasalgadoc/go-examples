package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Crear un canal de capacidad limitada
	c := make(chan int, 5)

	// Por cada iteración del for se usará un espacio del canal hasta llenarlo
	for i := 0; i < 10; i++ {
		wg.Add(1)
		c <- 1
		// Cuando se supere la capacidad del canal no podrá ejecutarse más GoRoutines hasta que haya espacio
		go doSomething(i, &wg, c)
	}

	fmt.Println("Waiting for goroutines to finish")

	// Aquí forzamos que tengan que ejecutarse las 10 GoRoutines del for a pesar de que por canales, no se hayan ocupado todas
	wg.Wait()

}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("id: %d started \n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id: %d finished \n", i)

	// Cuando finalice la función, podrá devolver la capacidad al canal
	<-c
}
