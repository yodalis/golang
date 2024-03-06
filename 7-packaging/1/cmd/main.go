// cmd é onde fica os arquivos executaveis
package main

import (
	"fmt"

	"github.com/yodalis/golang/7-packaging/1/math"
)

func main() {
	// m := math.Math{A: 1, B: 4}
	m := math.NewMath(5, 7)
	fmt.Println(m.Add())
	fmt.Println(m) // consegue ver os valores mas n tem acesso
}
