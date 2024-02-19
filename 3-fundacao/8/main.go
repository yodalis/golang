// Funcs
package main

func main() {
	sum(1, 2)
	println(sum(1, 2))
	println(sum2(25, 25))
}

// Forma normal da funcao
// func sum(a int, b int) int {
// 	return a + b
// }

// Caso as funcoes possuam argumentos do mesmo tipo voce pode apresenta-los da seguinte maneira
func sum(a, b int) int {
	return a + b
}

// Funcao com dois retornos
func sum2(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}
