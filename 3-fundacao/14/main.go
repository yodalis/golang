// Interfaces
package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	isActive bool
	Address
}

// Toda pessoa que utilizar o metodo citado, vai estar atrelada a struct Client-Addres
// interfaces no go so pode implementar metodos, e nao propriedades, essa e a diferenca entre interfaces e structs
type Person interface {
	inactiveClient()
}

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

func (c Client) inactiveClient() {
	c.isActive = false
	fmt.Printf("O cliente %s esta desativado", c.Name)
}

func Desactivation(person Person) {
	person.inactiveClient()
}

func main() {
	tha := Client{
		Name:     "Tha",
		Age:      22,
		isActive: true,
	}

	tha.City = "São Paulo"
	tha.Address.City = "São Paulo"
	Desactivation(tha)
}
