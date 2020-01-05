package main

import "fmt"

func main() {
	frutas := make(map[string]int, 2)
	frutas["mango"] = 5
	frutas["manzana"] = 10
	fmt.Println(frutas)
}
