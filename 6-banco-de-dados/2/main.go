package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model //Essa propriedade cria coisas na tabela padrão automaticamente
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" // Configurando dsn com padrao de data/hora
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name:  "Macbook",
	// 	Price: 1000000000,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Mouse", Price: 300},
	// 	{Name: "Keyboard", Price: 500},
	// }

	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Print(product)

	// select com where
	// db.First(&product, "name = ?", "Keyboard") // Retorna o primeiro registro que encontrar
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Print(product)
	// }

	// select all com limit
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Print(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 300).Find(&products)
	// for _, product := range products {
	// 	fmt.Print(product)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%m%").Find(&products)
	// for _, product := range products {
	// 	fmt.Print(product)
	// }

	// alter
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// delete
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

	// Soft delete: o dado nunca é realmente deletado, mas é indicado em qual momento foi excluído, como uma forma de lixeira

}
