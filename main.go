package main

import (
	"github.com/jhoguer/go-db-gorm/model"
	"github.com/jhoguer/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

	// storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")

	invoice := model.InvoiceHeader{
		Client: "Alvaro Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
}
