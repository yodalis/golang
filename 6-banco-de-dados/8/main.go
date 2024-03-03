package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // relação + nome da tabela intermediária
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	// é necessário ter essa notação nos dois campos que fazem essa relação many2many, no caso muitos produtos podem ter muitas categorias
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" // Configurando dsn com padrao de data/hora
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Oi"
	tx.Debug().Save(&c)
	tx.Commit() //liberando o acesso novamente
}

// Lock otimista
// Podemos pensar como um versionamento no git, é quando a linha que você tá alterando,
// pode ser alterada a qualquer momento por outra pessoa, mais comumente usado

// Lock pessimista
// É quando você bloqueia aquela linha para realizar alguma ação e enquanto você não desbloquear
// ninguém consegue alterar aquela linha, como se bloqueasse o dado.
