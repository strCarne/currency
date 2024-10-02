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

func MustGORM() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL must be set")
	}

	//nolint:exhaustruct
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	connPool, err := db.InitGORM(dsn, gormConfig)
	if err != nil {
		panic(err)
	}

	sqlDB, err := connPool.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(MaxIdleConns)
	sqlDB.SetMaxOpenConns(MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connPool
}
