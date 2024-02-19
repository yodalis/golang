// Packages/Modulos
package main

import (
	"fmt"

	"github.com/curso-go/matematica"
	"github.com/google/uuid"
)

func main() {
	sum := matematica.Sum(10, 20)
	fmt.Printf("Resultado: %v", sum)

	fmt.Println(matematica.A)

	mobi := matematica.Carro{Marca: "Mobi"}

	fmt.Printf("Meu carro eh um %s\n", mobi.Marca)
	fmt.Println(mobi.Andar())

	fmt.Println(uuid.Future)
}
