// Arrays
package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Tha"
	e float64 = 1
	f ID      = 3
)

// escopo global

func main() {
	var myArray [3]int // array de 3 posicoes com os valores do tipo int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	// Mostra o tamanho do meu array
	fmt.Println(len(myArray))

	// Pegando meu ultimo item
	fmt.Println(len(myArray) - 1)

	// Percorrendo o array
	for i, value := range myArray {
		fmt.Printf("O valor do indice %d e o valor e %d ", i, value)
	}
}
