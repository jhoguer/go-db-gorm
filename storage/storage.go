package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// _ "gorm.io/driver/mysql"
	// _ "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New created the connection with DB
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newMySQLDB create new connection to mySQL db
func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", "root:toor@tcp(localhost:3306)/godbORM?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}

// newPostgresDB create new connection to mySQL db
func newPostgresDB() {
	once.Do(func() {

		// Lo que este aqui solo se ejecutara 1 vez
		var err error
		db, err = gorm.Open("postgres", "postgres://jhoguer:jhon198615@localhost:5432/godbORM?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

// DB return a unique instace of db
func DB() *gorm.DB {
	return db
}

// Controlando los nulos string
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// Controlando los nulos Time
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// // DAOProduct factory of product.Storage
// func DAOProduct(driver Driver) (product.Storage, error) {
// 	switch driver {
// 	case Postgres:
// 		return newPsqlProduct(db), nil
// 	case MySQL:
// 		return newMySQLProduct(db), nil
// 	default:
// 		return nil, fmt.Errorf("Driver not implement")
// 	}
// }
