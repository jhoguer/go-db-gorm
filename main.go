package main

import (
	"fmt"

	"github.com/jhoguer/go-db-gorm/model"
	"github.com/jhoguer/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	// Creando migraciones
	// storage.DB().AutoMigrate(
	// 	&model.Product{},
	// 	&model.InvoiceHeader{},
	// 	&model.InvoiceItem{},
	// )

	// Creando registros
	// product1 := model.Product{
	// 	Name:  "Curso de Go",
	// 	Price: 120,
	// }

	// obs := "Testing with Go"
	// product2 := model.Product{
	// 	Name:         "Curso de Testing",
	// 	Price:        150,
	// 	Observations: &obs,
	// }

	// product3 := model.Product{
	// 	Name:  "Curso de Python",
	// 	Price: 200,
	// }

	// storage.DB().Create(&product1)
	// storage.DB().Create(&product2)
	// storage.DB().Create(&product3)

	// Leyendo multiples regsitros
	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}
}
