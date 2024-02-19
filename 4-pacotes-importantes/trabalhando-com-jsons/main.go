package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // Esses carinhas nas crases são como decorators/notation do Go
	// você pode colocar um nome diferente do que está na struct dentro das aspas que vai ser como ele vai imprimir na saída do json
	//  por exemplo: o Numero com n minusculo
	Saldo int `json:"s"`  // se colocar um - entre as aspas, vai omitir esse dado, n vai mostrar na saída
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}
	
	// Transformando em JSON
	res, err := json.Marshal(conta)
	if err != nil {
		fmt.Println("Erro ao gerar o JSON: ", err)
	}
	// Se não converter os bytes em string, vai imprimir em bytes
	fmt.Println(string(res))

	// Encoder: transforma em JSON e escreve em um outro lugar
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"n": 123,"s": 1000}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println("Erro ao converter o JSON: ", err)
	}
	println(contaX.Saldo)
}