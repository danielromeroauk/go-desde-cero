package main

import (
	"fmt"
	"time"
)

func main() {
	hora := time.Now().Hour()

	switch {
	case hora < 12:
		fmt.Println("¡Buenos días!")
	case hora < 17:
		fmt.Println("¡Buenas tardes!")
	default:
		fmt.Printf("¡Buenas noches!")
	}
}
