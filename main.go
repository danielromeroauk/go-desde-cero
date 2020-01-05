package main

import "fmt"

func main() {
	edades := make([]uint8, 0, 5)

	fmt.Println(len(edades))
	fmt.Println(cap(edades))
}
