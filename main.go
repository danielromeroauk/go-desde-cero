package main

import (
	"fmt"
)

func main() {
	sumatoria := 0
	for contador := 1; contador <= 5; contador++ {
		sumatoria += contador
		if contador == 3 {
			// continue
			break
		}
		// fmt.Println(contador)
	}

	fmt.Println(sumatoria)
}
