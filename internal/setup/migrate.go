package setup

import (
	"github.com/strCarne/currency/internal/schema"
	"gorm.io/gorm"
)

func MustMigrate(conn *gorm.DB) {
	//nolint:exhaustruct
	if err := conn.AutoMigrate(&schema.Rate{}); err != nil {
		panic("failed to migrate Rate")
	}
}
