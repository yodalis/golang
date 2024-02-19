// Slices
package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 100}

	fmt.Printf("len= %d capacidade=%d %v\n", len(s), cap(s), s)

	// o s[:0] indica que eu n quero que mostre nada da direita para a esquerda e mesmo que nâo mostre
	// isso não altera a capacidade inicial de 9 itens no slice
	fmt.Printf("len= %d capacidade=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Aqui ele vai mostrar os 4 primeiro valores e dropar os ultimos depois disso
	fmt.Printf("len= %d capacidade=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Aqui ele vai ignorar os dois primeiros valores e mostrar o resto
	// Como eu estou diminuindo do comeco pro fim, a capacidade do slice e diminuida
	fmt.Printf("len= %d capacidade=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Sempre que voce fizer um append, o slice vai entender que vc precisa de um array maior, e dobra o tamanho do array por tras
	// por isso a capacidade virou 18

	s = append(s, 110)
	fmt.Printf("len= %d capacidade=%d %v\n", len(s), cap(s), s)

}
