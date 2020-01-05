package main

import (
	"fmt"
)

func main() {
	idiomas := map[string]string{"es": "Español", "en": "Inglés"}

	for llave, valor := range idiomas {
		fmt.Printf("%s: %s\n", llave, valor)
	}
}
