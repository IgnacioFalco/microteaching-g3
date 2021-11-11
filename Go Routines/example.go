package main

import (
	"fmt"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

func main() {
	for i := 0; i < 10; i++ {
		proceso(i)
		//Mostrar como creariamos el proceso en una go routine
		//go proceso(i)
	}

	//Descomentar y Luego comentarlo para ver que sucede
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Termino el programa")

}

//Mas abajo esta el ejemplo de canales.

/*
func procesoCanal(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")

	c <- i
}

func main() {

	canal := make(chan int)

	go procesoCanal(1, canal)

	variable := <-canal // recibimos y lo asignamos a una variable

	fmt.Println("Termino el programa ", variable)

}
*/
