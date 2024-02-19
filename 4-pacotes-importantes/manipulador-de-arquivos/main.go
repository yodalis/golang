package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	f, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}
	// Se você sabe que o conteúdo do arquivo é uma string, 
	// você pode usar o método WriteString para escrever o conteúdo do arquivo de uma vez.

	// size, err := os.WriteString("Hello World")
	size, err := f.Write([]byte("Escrevendo no arquivo\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com %d bytes\n", size)

	f.Close()

	// leitura
	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo)) // Converte para string por que o conteudo dentro vem em bytes

	// leitura de pouco em pouco, abrindo o arquivo
	arquivo2, err := os.Open("test.txt") // Arquivo está aberto, mas não está lendo nada
	if err != nil {
		panic(err)
	}

	// Bufferizado = leitura de pouco em pouco
	reader := bufio.NewReader(arquivo2) // Reader é capaz de ler o arquivo
	// De quanto em quanto ele vai ler?
	// buffer := make([]byte, 10) // Vai ler de 10 em 10 bytes
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer) // lendo
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n])) // Converte e junta os pedaços para formar a string
	}

	// Removendo o arquivo
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}