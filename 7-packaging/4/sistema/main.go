package main

// Workspace local
// go work init + pastas que vou usar ex: go work init ./math ./sistema
//  para conseguir fazer go mod tidy sem problemas por causa do workspace é só fazer um
// go mod tidy -e - esse comando ignora os pacotes que ele n conseguiu achar

import (
	"github.com/google/uuid"
	"github.com/yodalis/golang/7-packaging/3/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
