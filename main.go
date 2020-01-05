package main

import "fmt"

func main() {
	nombre := "Daniel"
	apellido := "Romero"
	edad := 28

	fmt.Printf("nombre es %T, apellido es %T y edad es %T\n", nombre, apellido, edad)
	fmt.Printf("nombre=%s\n", nombre)
	fmt.Printf("apellido=%s\n", apellido)
	fmt.Printf("edad=%d\n", edad)
}
