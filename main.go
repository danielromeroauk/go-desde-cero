package main

import (
	"fmt"
	"reflect"
)

func main() {
	// personas := []string{"María", "Hernan", "Olga"}
	// saludar(personas...)
	fmt.Println(sumatoria(5, 7, 1, 8, 3))
}

func saludar(nombres ...string) {
	fmt.Println(reflect.TypeOf(nombres))

	for _, nombre := range nombres {
		fmt.Println("Hola", nombre)
	}
}

func sumatoria(numeros ...int) int {
	var sumatoria int
	for _, numero := range numeros {
		sumatoria += numero
	}

	return sumatoria
}
