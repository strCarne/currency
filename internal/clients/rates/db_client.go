package rates

import (
	"context"

	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/pkg/models"
)

type DBClientSelect interface {
	SelectRates(ctx context.Context) ([]schema.Rate, error)
	SelectRateByID(ctx context.Context, id int) (*schema.Rate, error)
	SelectRatesByCurID(ctx context.Context, curID int) ([]schema.Rate, error)
	SelectRatesByCurIDAndDate(ctx context.Context, curID int, onDate models.Date) ([]schema.Rate, error)
}

type DBClientInsert interface {
	InsertRate(ctx context.Context, rate *schema.Rate) error
	InsertRates(ctx context.Context, rates []schema.Rate) error
}

type DBClient interface {
	DBClientSelect
	DBClientInsert
}
