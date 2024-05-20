package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yodalis/golang/10-multithreading/desafio/internal/entity"
	"github.com/yodalis/golang/10-multithreading/desafio/internal/infra/api"
)

type AddressHandler struct {
	ViaCEP    api.ViaCEPInterface
	BrasilCEP api.BrasilAPIInterface
}

func NewAddressHandler(viaCep api.ViaCEPInterface, brasilCep api.BrasilAPIInterface) *AddressHandler {
	return &AddressHandler{
		ViaCEP:    viaCep,
		BrasilCEP: brasilCep,
	}
}

func (h *AddressHandler) GetAddressByCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := h.getAddressResponse(cep)
	if err != nil {
		if err.Error() == "Timeout exceeded" {
			w.WriteHeader(http.StatusGatewayTimeout)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func (h *AddressHandler) getAddressResponse(cep string) (*entity.Address, error) {
	cViaCep := make(chan *entity.Address)
	cBrasilApi := make(chan *entity.Address)

	go func() {
		address, _ := h.ViaCEP.GetAddressByCep(cep)
		cViaCep <- address
	}()
	go func() {
		address, _ := h.BrasilCEP.GetAddressByCep(cep)
		cBrasilApi <- address
	}()

	select {
	case address := <-cViaCep:
		fmt.Printf("Received from ViaCEP:\n CEP: %s\n StreetName: %s\n City: %s\n State: %s\n Neighborhood: %s\n", address.CEP, address.StreetName, address.City, address.State, address.Neighborhood)
		return address, nil
	case address := <-cBrasilApi:
		fmt.Printf("Received from Brasil API:\n CEP: %s\n StreetName: %s\n City: %s\n State: %s\n Neighborhood: %s\n", address.CEP, address.StreetName, address.City, address.State, address.Neighborhood)
		return address, nil
	case <-time.After(time.Second):
		fmt.Println("Timeout exceeded")
		err := errors.New("Timeout exceeded")
		return nil, err
	}
}
