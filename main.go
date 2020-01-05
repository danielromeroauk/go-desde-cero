package main

import "fmt"

func main() {
	var edades [3]uint8
	edades[0] = 18
	edades[1] = 20
	edades[2] = 30

	for indice, valor := range edades {
		fmt.Printf("En el índice %d está el valor %d\n", indice, valor)
	}
}
