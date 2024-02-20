package main

import (
	"database/sql"
	"fmt"

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

	// p, err := selectOneProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
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

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	//Query Row busca apenas uma linha
	// Scan vai olhar todas as colunas, encontrar as correspondentes e então preencher as colunas
	// do p com os valores encontrados correspondentes
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Sem necessidade de proteger da sql injection
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		// Criando cada product encontrado em um Product (struct)
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		// Jogando cada Product dentro do slice que será devolvido
		products = append(products, p) // products = lista de products + p
	}

	return products, nil
}
