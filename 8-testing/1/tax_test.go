package tax

import "testing"

// Para testar -> go test .
// Para testar de forma verbosa -> go test -v
// Como pegar o coverage -> go test -coverprofile=coverage.out
// Como saber onde não está coberto -> go tool cover -html=coverage.out
// Vai indicar no navegador o ponto que n está coberto
// testing.T -> de test

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

// Batch -> calculando varios de uma vez
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

// Benchmark = refere-se a um tipo específico de teste utilizado para medir o desempenho de uma função ou de um trecho de código em relação a um ou mais parâmetros de referência
// BenchmarkCalculateTax-12        1000000000               0.7232 ns/op
// -12 é a quantidade de nucleos computacionais utilizados para executar a func
// 1000000000 quantidade de operações que foram executadas
//  0.7232 ns/op quantas operações por nano segundo
// comando para rodar os testes + benchmark =  go test -bench .
// comando pra rodar apenas o benchmark = go test -bench=. -run=^#
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// Verificar qual dos dois algoritmos term a melhor performance
// testing.B de benchmark
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// Faz 10 vezes o benchmark = go test -bench=. -run=^# -count=10
// Quanto trabalhando na alocação de memória = go test -bench=. -run=^# -benchmem

// Fuzzing = testes de mutação
// Gerador de possibilidades para meus testes
// testing.F de fuzz
// go test -fuzz=.
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
	 f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
	 result := CalculateTax(amount)
	 if amount <= 0 && result != 0 {
		 t.Errorf("Reveived %f but expected 0", result)
	 }
	 if amount > 20000 && result != 20 {
		 t.Errorf("Reveived %f but expected 20", result)
	 }
	})
 }
