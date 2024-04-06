package main

import (
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yodalis/golang/9-apis/configs"
	"github.com/yodalis/golang/9-apis/internal/entity"
	"github.com/yodalis/golang/9-apis/internal/infra/database"
	"github.com/yodalis/golang/9-apis/internal/infra/webserver/handlers"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetAllProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8000", r)
}
