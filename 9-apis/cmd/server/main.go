package main

import (
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/yodalis/golang/9-apis/configs"
	"github.com/yodalis/golang/9-apis/internal/entity"
	"github.com/yodalis/golang/9-apis/internal/infra/database"
	"github.com/yodalis/golang/9-apis/internal/infra/webserver/handlers"
)

func main() {
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	// Guardando o valor do token auth no context do middleware de jwt do chi para usar como uma espécie de variável
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))

	// Agrupando as rotas de Products
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth)) //Middleware que é usado para verificar a autenticidade de tokens JWT em requisições HTTP
		r.Use(jwtauth.Authenticator)               //Esse middleware é usado para extrair e analisar tokens JWT de uma solicitação HTTP
		//  e colocar as claims (infos do usuario que são codificadas) do token em contexto.
		// Claims Registradas -> São claims definidas pela especificação JWT e têm um significado pré-definido.
		// Algumas das claims registradas incluem iss (emissor), sub (assunto), aud (público-alvo), exp (data de expiração) e iat (data de emissão).
		// Que é a que estamos usando aqui nessa api
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetAllProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}
