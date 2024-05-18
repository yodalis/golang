package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yodalis/golang/10-multithreading/desafio/internal/entity"
)

type ViaCEP struct {
}

type ViaCEPRes struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func NewViaCEP() *ViaCEP {
	return &ViaCEP{}
}

func (h *ViaCEP) GetAddressByCep(cep string) (*entity.Address, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var viaCEPData *ViaCEPRes
	json.Unmarshal(responseData, &viaCEPData)

	address := entity.NewAddress(viaCEPData.CEP, viaCEPData.Logradouro, viaCEPData.Localidade, viaCEPData.UF, viaCEPData.Bairro)

	return address, nil
}
