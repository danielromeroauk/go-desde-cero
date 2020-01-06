package main

import (
	"fmt"
	"log"
)

func main() {
	miFuncion := saludar

	saludo, _, err := miFuncion("Daniel", "Romero", 19)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("saludo:", saludo)
}

func saludar(nombre, apellido string, edad uint8) (texto string, booleano bool, err error) {
	booleano = edad >= 18
	texto = fmt.Sprintf("Hola %s %s", nombre, apellido)

	if len(apellido) <= 2 {
		err = fmt.Errorf("El apellido '%s' no es válido", apellido)
	}

	return
}
