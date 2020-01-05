package main

import "fmt"

func main() {
	edades := [3]uint8{18, 20, 30}

	for indice, valor := range edades {
		fmt.Printf("En el índice %d está el valor %d\n", indice, valor)
	}
}
