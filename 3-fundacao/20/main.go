// Generics - Inicial
package main

// Problematica: Criar duas funções diferentes, só por causa do tipo delas

// func somaInteiro(m map[string]int) int {
// 	var soma int

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }
// func somaFloat(m map[string]float64) float64 {
// 	var soma float64

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// Resolução: colocando o T dessa forma você indica que pode ser tanto int quanto float64 ouuuuu

// func Soma[T int | float64](m map[string]T) T {
// 	var soma T

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// Resolução 2: Você pode criar uma constrain que vai indicar isso mas de forma mais simples

type MyNumber int

type Number interface {
	~int | ~float64 // Aqui você indica uma exceção em que como o MyNumber tem um int, ele passa mesmo sendo de tipos diferentes
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a T, b T) bool { // Comparable é uma constrain que permite apenas valores que possam
	//  ser comparados, ex: int com int, PORÉM o comparable não supre casos como a > b, por que ele não sabe se isso pode acontecer
	// Exemplo: uma string pode ser igual outra, porém não tem como ela ser maior do que outra numericamente falando,
	// então ele n supre esse ponto

	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Godao": 200, "Kiri": 300, "Pitas": 500}
	m2 := map[string]float64{"Godao": 200.2, "Kiri": 300.4, "Pitas": 500.6}
	m3 := map[string]MyNumber{"Godao": 100, "Kiri": 500, "Pitas": 900}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 20))
}
