package main

import (
	"fmt"
)

func main() {
	saludo := saludar("Daniel")

	fmt.Println(saludo)
}

func saludar(nombre string) string {
	return fmt.Sprintf("Hola %s", nombre)
}
