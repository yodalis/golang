// Composicao de Structs

package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	isActive bool
	Address  // Design Pattern utilizado aqui e a Composicao, se criasse com um nome diferente,
	//  seria uma inferenca de tipo a variavel criada
}

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

func main() {
	tha := Client{
		Name:     "Tha",
		Age:      22,
		isActive: true,
	}

	// 2 formas de adicionar valor a struct, a primeira ja conseguimos direto pois estamos compondo uma struct com outra
	tha.City = "São Paulo"
	tha.Address.City = "São Paulo"

	fmt.Printf("Name: %s - Age: %d - isActive: %t", tha.Name, tha.Age, tha.isActive)
}
