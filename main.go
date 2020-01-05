package main

import "fmt"

func main() {
	if edad := 15; edad >= 18 {
		fmt.Println("Eres mayor de edad en Colombia")
		fmt.Println(edad)
	} else {
		fmt.Println("Eres menor de edad en Colombia")
		fmt.Println(edad)
	}
}
