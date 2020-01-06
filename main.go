package main

import (
	"fmt"
	"reflect"
)

func main() {
	saludar("María", "Hernan", "Olga")
}

func saludar(nombres ...string) {
	fmt.Println(reflect.TypeOf(nombres))

	for _, nombre := range nombres {
		fmt.Println("Hola", nombre)
	}
}
