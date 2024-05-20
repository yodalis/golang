package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yodalis/golang/10-multithreading/desafio/internal/infra/api"
)

type AddressHandler struct {
	ViaCEP api.ViaCEPInterface
}

func NewAddressHandler(externalApi api.ViaCEPInterface) *AddressHandler {
	return &AddressHandler{
		ViaCEP: externalApi,
	}
}

func (h *AddressHandler) GetAddressByCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := h.ViaCEP.GetAddressByCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}
