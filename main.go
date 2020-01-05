package main

import (
	"fmt"
)

func main() {
	nombre := "Daniel"

	// for posicion, valor := range nombre {
	// 	// fmt.Printf("Posición: %T, Valor: %T\n", posicion, valor)
	// 	fmt.Printf("Posición: %d, Valor: %s\n", posicion, string(valor))
	// }

	// for _, valor := range nombre {
	// 	// fmt.Printf("Posición: %T, Valor: %T\n", posicion, valor)
	// 	fmt.Printf("Valor: %s\n", string(valor))
	// }

	for _, _ = range nombre {
		// fmt.Printf("Posición: %T, Valor: %T\n", posicion, valor)
		fmt.Printf("Hola\n")
	}
}
