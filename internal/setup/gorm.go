package setup

import (
	"os"
	"time"

	"github.com/strCarne/currency/pkg/db"
	"gorm.io/gorm"
)

const (
	MaxIdleConns = 10
	MaxOpenConns = 100
)

func MustGORM() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL must be set")
	}

	//nolint:exhaustruct
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	if err := db.InitGORM(dsn, gormConfig); err != nil {
		panic(err)
	}

	pool, err := db.Connection().DB()
	if err != nil {
		panic(err)
	}

	pool.SetMaxIdleConns(MaxIdleConns)
	pool.SetMaxOpenConns(MaxOpenConns)
	pool.SetConnMaxLifetime(time.Hour)
}
