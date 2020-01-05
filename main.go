package main

import (
	"fmt"
)

func main() {
	animales := [2]string{"perro", "gato"}

	for _, valor := range animales {
		fmt.Println(valor)
	}
}
