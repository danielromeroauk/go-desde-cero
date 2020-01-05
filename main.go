package main

import (
	"fmt"
)

func main() {
	saludo := saludar()
	fmt.Println(saludo)
}

func saludar() string {
	return "Hola desde una función"
}
