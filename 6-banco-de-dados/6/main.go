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
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model   //Essa propriedade cria coisas na tabela padrão automaticamente
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" // Configurando dsn com padrao de data/hora
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// create product
	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      50.00,
	// 	CategoryID: 1,
	// })

	// create serial number - has one
	// db.Create(&SerialNumber{
	// 	Number:    "1234",
	// 	ProductID: 1,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "Serial Number: ", product.SerialNumber.Number)
		}
	}

}
