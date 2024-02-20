package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	// isso é necessário pq a linguagem n escolhe qual driver você vai user
	// driver = oracle, mysql, postgrees
	// _ vai deixar a gente não utilizar diretamente o pacote, mesmo que seja necessário, sem isso o go n deixa a gente compilar
	// o programa sem usar o pacote diretamente

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// formação do segundo parâmetro: user:password@tipoderede(url)/banco
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 2000)

	err = insertProduct(db, product)

	if err != nil {
		panic(err)
	}
	product.Price = 3000
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	// Evitando sql injection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}
