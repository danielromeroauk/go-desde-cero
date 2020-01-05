package main

import "fmt"

func main() {
	frutas := make(map[string]int, 2)
	frutas["mango"] = 5
	frutas["manzana"] = 10

	for llave, valor := range frutas {
		fmt.Printf("Llave %s: %d\n", llave, valor)
	}
}
