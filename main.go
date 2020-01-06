package main

import (
	"fmt"
)

func main() {
	saludo, mayor := saludar("Daniel", "Romero", 19)

	fmt.Println("saludo:", saludo)
	fmt.Println("mayor:", mayor)
}

func saludar(nombre, apellido string, edad uint8) (texto string, booleano bool) {
	booleano = edad >= 18
	texto = fmt.Sprintf("Hola %s %s", nombre, apellido)

	return
}
