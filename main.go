package main

import (
	"fmt"
)

func main() {
	frutas := []string{"manzana", "pera", "mango"}

	for _, valor := range frutas {
		fmt.Println(valor)
	}
}
