package matematica

func Sum[T int | float64](a T, b T) T {
	return a + b
}

var A int = 22

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}
