package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	// Canal con dos mensajes
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// La longitud refleja los canales asignamdos y el cap la capacidad real del canal
	fmt.Println(len(c), cap(c))

	// Cerrar el canal: No se reciben datos adicionales, haga esto para evitar errores
	// al ejecutar range
	close(c)

	// Recorrer los datos de un canal
	for message := range c {
		fmt.Println(message)
	}

	// Select: En Multiples canales para obtener un canal

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	// ¿Cuál canal respondería primero?
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}

}
