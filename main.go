package main

import (
	"github.com/jhoguer/go-db-gorm/model"
	"github.com/jhoguer/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Creando migraciones
	// storage.DB().AutoMigrate(
	// 	&model.Product{},
	// 	&model.InvoiceHeader{},
	// 	&model.InvoiceItem{},
	// )

	product1 := model.Product{
		Name:  "Curso de Go",
		Price: 120,
	}

	obs := "Testing with Go"
	product2 := model.Product{
		Name:         "Curso de Testing",
		Price:        150,
		Observations: &obs,
	}

	product3 := model.Product{
		Name:  "Curso de Python",
		Price: 200,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
}
