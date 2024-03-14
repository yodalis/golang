package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax) //Informa q a func será chamada
	return args.Error(0)  // Aqui é o tipo de retorno da sua func, se fosse um int seria args.Int
}
