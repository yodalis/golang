package main

const a = "Hello World"

var (
	b bool
	c int
	d string
	e float64
)

// escopo global

func main() {
	var a string // escopo local
	println(a)
}
