// Ponteiros / Structs
package main

import "fmt"

// Exemplo 1

// type Client struct {
// 	Name string
// }

// func (c Client) walk() {
// 	c.Name = "Tha Calazans"
// 	fmt.Printf("O cliente %s andou\n", c.Name)
// }

// func main() {
// 	tha := Client{
// 		Name: "tha",
// 	}
// 	tha.walk()

// 	fmt.Printf("O valor do nome do client é %s", tha.Name)
// }

// Exemplo 2
type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{Saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.Saldo += valor

	return c.Saldo
}

func main() {
	conta := Conta{
		Saldo: 100,
	}

	conta.simular(500)
	fmt.Printf("O novo valor da conta é: %d", conta.Saldo)
}
