package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yodalis/golang/10-multithreading/desafio/internal/entity"
)

type BrasilCEPApi struct {
}

type BrasilCepRes struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilCep() *BrasilCEPApi {
	return &BrasilCEPApi{}
}

func (h *BrasilCEPApi) GetAddressByCep(cep string) (*entity.Address, error) {
	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cepData *BrasilCepRes
	json.Unmarshal(responseData, &cepData)

	address := entity.NewAddress(cepData.CEP, cepData.Street, cepData.City, cepData.State, cepData.Neighborhood)

	return address, err
}
