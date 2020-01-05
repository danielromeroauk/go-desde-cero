package main

import (
	"fmt"
)

func main() {
	contador := 1

	// while
	// for contador < 5 {
	// 	fmt.Println(contador)
	// 	contador++
	// }

	// do while
	for {
		fmt.Println(contador)

		if contador >= 5 {
			break
		}
		contador++
	}

	fmt.Println("Fin.")
}
