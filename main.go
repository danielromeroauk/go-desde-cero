package main

import "fmt"

func main() {
	var (
		nombre   string
		apellido string
	)
	nombre = "Daniel"
	apellido = "Romero"
	fmt.Printf("nombre es %T y apellido es %T\n", nombre, apellido)
	fmt.Printf("nombre=%s\n", nombre)
	fmt.Printf("apellido=%s\n", apellido)
}
