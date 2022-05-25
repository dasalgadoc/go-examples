package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	// Al final de la ejecución de esta función, notifica el término.
	defer wg.Done()

	fmt.Println(text)
}

func main() {

	// Creación del Grupo de Sincronización
	var wg sync.WaitGroup

	fmt.Println("Hola")

	// Agregar el GoRoutine al Grupo (Función Say)
	wg.Add(1)

	// Enviar el grupo como parámetro con el apuntador
	go say("Mundo", &wg)

	// Espera a que finalice todas las GoRoutines
	wg.Wait()

	// Esta línea ya no es necesaria
	//time.Sleep(time.Second * 1)
}
