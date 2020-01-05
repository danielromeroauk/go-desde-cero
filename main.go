package main

import "fmt"

func main() {
	frutas := make([]string, 0, 10)
	frutas = append(frutas, "mango")

	fmt.Println(frutas)
}
