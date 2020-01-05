package main

import "fmt"

func main() {
	frutas := make([]string, 0, 10)
	frutas = frutas[0:6]
	frutas[5] = "mango"

	fmt.Println(frutas)
}
