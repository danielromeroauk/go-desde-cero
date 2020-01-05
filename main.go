package main

import (
	"fmt"
)

func main() {
	saludo := saludar("Daniel", "Romero")

	fmt.Println(saludo)
}

func saludar(nombre string, apellido string) string {
	return fmt.Sprintf("Hola %s %s", nombre, apellido)
}
