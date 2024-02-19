// Maps
package main

import "fmt"

func main() {
	salarios := map[string]int{"Tha": 200, "Be": 400, "Mari": 700}

	// Apagando valores do map
	delete(salarios, "Tha")

	// Adicionando valores ao map
	salarios["Thais"] = 5000
	fmt.Println(salarios["Thais"])

	// Func make -> forma de criar um map sem inicializar
	// sal := make(map[string]int)

	// Percorrendo maps
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s e %d \n", nome, salario)
	}
	// Ignorando o nome
	for _, salario := range salarios {
		fmt.Printf("O salario e %d \n", salario)
	}
}
