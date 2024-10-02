package db

import (
	"errors"
	"os"

	"github.com/strCarne/currency/pkg/wrapper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrInitGORM = errors.New("filed to initialize GORM")

func InitDefaultGORM() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, ErrInitGORM
	}

	//nolint:exhaustruct
	connPool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, wrapper.Wrap("db.gorm.InitDefaultGORM", "failed to open", err)
	}

	return connPool, nil
}

func InitGORM(dsn string, config *gorm.Config) (*gorm.DB, error) {
	connPool, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, wrapper.Wrap("db.gorm.InitDefaultGORM", "failed to open", err)
	}

	return connPool, nil
}
