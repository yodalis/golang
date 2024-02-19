package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
}

func main() {
		for _, cep := range os.Args[1:] {
			req, err := http.Get("http://viacep.com.br/ws/"+ cep + "/json/")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err) // O Fprintf escreve no stderr
			}
			defer req.Body.Close()

			res, err := io.ReadAll(req.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
			}

			var data ViaCEP
			err = json.Unmarshal(res, &data)  // Transformando o json em struct
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao converter a resposta: %v\n", err)
			}

			fmt.Println(data.Bairro)

			file, err := os.Create("bairro.txt")
			if err != nil {	
				fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
			}

			defer file.Close()

			_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))	// o Sprintf joga a saída para algum lugar
			fmt.Println("Arquivo criado com sucesso")
		}
}