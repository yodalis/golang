package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model //Essa propriedade cria coisas na tabela padrão automaticamente
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" // Configurando dsn com padrao de data/hora
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Panela",
		Price:      50.00,
		CategoryID: 5,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}

}
