package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("saludo1:", reflect.TypeOf(saludo1))
	fmt.Println("saludo2:", reflect.TypeOf(saludo2))

	conNombre := saludo1("Daniel")
	conApellido := saludo2("Guillermo")

	fmt.Println(conNombre)
	fmt.Println(conApellido)
}

func saludo1(nombre string) string {
	return fmt.Sprintf("Hola %s", nombre)
}

func saludo2(nombre string) (texto string) {
	texto = fmt.Sprintf("Hola %s", nombre)
	return
}
