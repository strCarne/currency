package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SAFETY:
// gorm.DB is thread-safe, so it is ok to use it in multiple goroutines
// and be a global variable.
//
//nolint:gochecknoglobals
var connection *gorm.DB

func initGORM() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set")
	}

	//nolint:exhaustruct
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection = conn
}

func Connection() *gorm.DB {
	if connection == nil {
		initGORM()
	}

	return connection
}
