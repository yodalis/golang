// Type assertion
// Conversão de tipos
package main

import "fmt"

func main() {
	var myVar interface{} = "Tha Calazans"
	res, ok := myVar.(int) // O ok vai indicar se é possível de realizar a conversão

	fmt.Printf("O valor de res é: %v e o resultado de ok é - %v", res, ok)

	res2 := myVar.(int)
	fmt.Printf("O valor de res é: %v ", res2) // Da um panic pq ele não consegue verificar se realmente n da
}
