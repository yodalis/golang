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
	// T mostra o tipo da variavel
	fmt.Printf("O tipo de E e %T", e)
}
