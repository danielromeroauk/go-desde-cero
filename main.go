package main

import "fmt"

import "reflect"

func main() {
	var nombre string
	// nombre = "Daniel" + " Romero"
	// nombre = "Daniel" + "\nRomero"
	// nombre = "Daniel\nRomero"
	nombre = `Daniel Romero`

	fmt.Println(reflect.TypeOf(nombre[0]))

	// fmt.Println(nombre[0])
	// fmt.Println(string(nombre[0]))

	for posicion, letra := range nombre {
		fmt.Printf("Posición %d, letra \"%s\"\n", posicion, string(letra))
	}

	fmt.Printf("Hay %d letras en total\n", len(nombre))
}
