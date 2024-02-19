// Closures
// funcoes dentro de funcoes - funcoes anonimas

package main

func main() {
	// Aqui estamos atribuindo o valor da variavel como uma funcao anonima (que pode ser considerada uma rotina)
	// que vai retornar o valor da func sum multiplicado por 2 e depois vamos printar isso.
	// Essas funcoes normalmente sao utilizadas para criar rotinas, como inicializar servicos de forma anonima
	total := func() int {
		return sum(1, 2, 9, 9, 902, 332, 5) * 2
	}()

	println(total)
}

// Caso voce n saiba a quantidade de parametros que sera passado, voce pode utilizar o ...
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
