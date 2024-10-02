package db

import (
	"errors"
	"os"

	"github.com/strCarne/currency/pkg/wrapper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SAFETY:
// gorm.DB is thread-safe, so it is ok to use it in multiple goroutines
// and be a global variable.
//
//nolint:gochecknoglobals
var connection *gorm.DB

var (
	ErrInitGORM               = errors.New("filed to initialize GORM")
	ErrGORMAlreadyInitialized = errors.New("GORM already initialized")
)

func InitDefaultGORM() error {
	if connection != nil {
		return ErrGORMAlreadyInitialized
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return ErrInitGORM
	}

	//nolint:exhaustruct
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return wrapper.Wrap("db.gorm.InitDefaultGORM", "failed to open", err)
	}

	connection = conn

	return nil
}

func InitGORM(dsn string, config *gorm.Config) error {
	if connection != nil {
		return ErrGORMAlreadyInitialized
	}

	conn, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return wrapper.Wrap("db.gorm.InitDefaultGORM", "failed to open", err)
	}

	connection = conn

	return nil
}

func Connection() *gorm.DB {
	if connection == nil {
		if err := InitDefaultGORM(); err != nil {
			panic(err)
		}
	}

	return connection
}
