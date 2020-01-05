package main

import "fmt"

func main() {
	edades := make([]uint8, 0, 10)
	edades = edades[0:6]
	edades[5] = 28

	fmt.Println(edades)
}
