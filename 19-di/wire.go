//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/yodalis/golang/19-di/product"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	// Indicando quais são os módulos que utilizo, e ele adivinha qual a ordem de execução de dependencia deles
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}
