package main

import "github.com/jhoguer/go-db-gorm/storage"

func main() {
	driver := storage.MySQL
	storage.New(driver)
	storage.New(driver)
	storage.New(driver)
	storage.New(driver)
	storage.New(driver)
	storage.New(driver)
}
