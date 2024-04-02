package database

import (
	"github.com/yodalis/golang/9-apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product

	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}

// Paginação -> O limit é a quantidade de dados que vai trazer, o offset é qual pagina você está
// Offset = 0 são as primeiras paginas igual um array

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		// Aqui estamos dizendo, o limite determinado é o limite do parâmetro, o offset será da pagina menos 1 (para n existir pag 0)
		// * limit para saber a partir de qual pagina ele vai e o Order é que vai ser ordenado pelos created_at de form crescente (asc)
		// Encontra os valores com find e preencha a var products com o que foi encontrado
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at").Find(&products).Error
	} else {
		err = p.DB.Order("created_at").Find(&products).Error
	}

	return products, err
}
