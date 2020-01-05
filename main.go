package main

import "fmt"

import "runtime"

func main() {
	fmt.Print("GO está corriendo en ")

	switch sistemaOperativo := runtime.GOOS; sistemaOperativo {
	case "darwin":
		fmt.Println("macOS.")

		// ejecuta el cuerpo del caso siguiente sin evaluarlo
		fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", sistemaOperativo)
	}
}
