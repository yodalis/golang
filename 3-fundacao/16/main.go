// Quando usar ponteiros
package main

// func soma(a, b int) int {
// 	a = 50 // Copia da variavel, será afetado apenas nesse escopo
// 	return a + b
// }

func soma(a, b *int) int {
	*a = 50
	*b = 150
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20

	soma(&var1, &var2)

	println(var1) // alterando os valores direto na memória
	println(var2) // alterando os valores direto na memória
	println(soma(&var1, &var2))
}
