package api

import "github.com/yodalis/golang/10-multithreading/desafio/internal/entity"

type BrasilAPIInterface interface {
	GetAddressByCep(cep string) (*entity.Address, error)
}

type ViaCEPInterface interface {
	GetAddressByCep(cep string) (*entity.Address, error)
}
