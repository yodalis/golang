// Funcs variadicas
package main

func main() {
	sum(1, 2)
	println(sum(1, 2, 9, 9, 902, 332, 5))
}

// Caso voce n saiba a quantidade de parametros que sera passado, voce pode utilizar o ...
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
