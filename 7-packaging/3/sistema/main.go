package main

import "github.com/yodalis/golang/7-packaging/3/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
