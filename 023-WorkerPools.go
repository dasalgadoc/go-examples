package main

import "fmt"

// Recibe dos canales, entrada y salida. Datos compartidos para los Workers
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("--> Worker #%d started fib with %d\n", id, job)
		fib := Fibo(job)
		fmt.Printf("Worker #%d finished. Job: %d; Fibo: %d\n", id, job, fib)
		results <- fib
	}
}

func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}

func main() {
	// Lista de tareas
	tasks := []int{2, 3, 4, 7, 10, 15}

	// Número de Workers
	nWorkers := 3

	// Canales de Enntrada y salida de la longitud de las tareas a ejecutar
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// Se crean los 3 Workers concurrentes y cada uno inicia su ejecución independiente, leyendo del canal compartido
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	// El canal de entrada se llena con las tareas
	for _, v := range tasks {
		jobs <- v
	}

	// Todas las tareas fueron ingresadas
	close(jobs)

	// Espere la lectura de 10 resultados
	for i := 0; i < len(tasks); i++ {
		<-results
	}
}
