// Structs
// Go nao e orientado a objetos, tem um "go way" de se programar
package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	isActive bool
}

func main() {
	tha := Client{
		Name:     "Tha",
		Age:      22,
		isActive: true,
	}

	fmt.Printf("Name: %s - Age: %d - isActive: %t", tha.Name, tha.Age, tha.isActive)
}
