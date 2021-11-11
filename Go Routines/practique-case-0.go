package main

import (
	"fmt"
	"math/rand"
	"time"
)

func procesoQuePuedeTardarMucho(nombre string) int {

	inicio := time.Now().UnixMilli() // inicio y fin son usados para medir el tiempo de ejecucion
	// aca simulamos un proceso que puede ser lento y de tiempo variable
	// como una consulta HTTP a una API, un query a una base de datos
	// lectura de un archivo en disco, etc
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5*1000))) // esperamos un tiempo aleatorio, entre 0 y 5 segundos
	r := rand.Intn(100)
	fin := time.Now().UnixMilli()
	fmt.Printf("%s duró %d milisegundos\n", nombre, fin-inicio)

	return r

}

func sumar(a, b int) int {
	fmt.Printf("a vale %d, y b vale %d\n", a, b)
	return a + b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	inicio := time.Now().UnixMilli()
	a, b := procesoQuePuedeTardarMucho("Tarea A"), procesoQuePuedeTardarMucho("Tarea B")
	fmt.Printf("El resultado de la suma es: %d\n", sumar(a, b))
	fin := time.Now().UnixMilli()
	fmt.Printf("Toda la ejecucion duró %d milisegundos\n", fin-inicio)
}
