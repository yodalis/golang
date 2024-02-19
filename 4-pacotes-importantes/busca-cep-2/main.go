package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", BuscaCEPHandler) // Recebe uma função que recebe um http.ResponseWriter e um http.Request
	// Poderia colocar a função diretamente no segundo parâmetro da http.HandleFunc
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) { // w é o que vamos responder e r é a requisição
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) // Escreve o status da resposta
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Caso fizessemos com marshal ficaria assim:
	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	json.NewEncoder(w).Encode(cep) // Pega a response e coloca no w que é o writer da resposta, dessa forma não tem necessidade de dar um marshal
	// Caso queira armazenar, usa um Marshal
}

func BuscaCep(cep string) (*ViaCEP, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
