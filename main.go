package main

import "fmt"

import "reflect"

func main() {
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(5 < 10 && 5 != 50) // true
}
