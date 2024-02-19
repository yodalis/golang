// Ponteiros
package main

func main() {
	// Memoria -> Endereco -> Valor

	a := 10
	// Valor do a
	println(a)

	// qual o endereco do valor de a
	println(&a)

	// mudando o valor de a pelo endereco de memoria
	var ponteiro *int = &a
	*ponteiro = 20

	// aqui estamos criando a variavel utilizando o endereco de memoria de a,
	// para que possamos mudar o valor direto no endereco de a, se voce aponta paara a, vc ve o valor que colocamos no ponteiro
	// por que mudamos esse valor direto no endereco do a
	println(a)

	b := &a
	// se printarmos dessa forma vai aparecer apenas o endereco e n o valor
	println(b)
	// qual valor guardado no b
	println(*b)
}
