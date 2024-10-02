package rates

import (
	"context"

	"github.com/strCarne/currency/internal/schema"
	"github.com/strCarne/currency/pkg/models"
	"gorm.io/gorm"
)

type DBClientStd struct {
	connPool *gorm.DB
}

func NewStd(connPool *gorm.DB) DBClientStd {
	return DBClientStd{
		connPool: connPool,
	}
}

func (d DBClientStd) SelectRates(ctx context.Context) ([]schema.Rate, error) {
	var rates []schema.Rate
	result := d.connPool.WithContext(ctx).Find(&rates)

	if result.Error != nil {
		return nil, result.Error
	}

	return rates, nil
}

func (d DBClientStd) SelectRateByID(ctx context.Context, id int) (*schema.Rate, error) {
	var rate schema.Rate
	result := d.connPool.WithContext(ctx).First(&rate, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rate, nil
}

func (d DBClientStd) SelectRatesByCurID(ctx context.Context, curID int) ([]schema.Rate, error) {
	var rates []schema.Rate
	result := d.connPool.WithContext(ctx).Where("currency_id = ?", curID).Find(&rates)

	if result.Error != nil {
		return nil, result.Error
	}

	return rates, nil
}

func (d DBClientStd) SelectRateByDate(ctx context.Context, onDate models.Date) ([]schema.Rate, error) {
	var rates []schema.Rate
	result := d.connPool.WithContext(ctx).Where("date = ?", onDate.String()).Find(&rates)

	if result.Error != nil {
		return nil, result.Error
	}

	return rates, nil
}

func (d DBClientStd) SelectRatesByCurIDAndDate(
	ctx context.Context,
	curID int,
	onDate models.Date,
) ([]schema.Rate, error) {
	var rates []schema.Rate
	result := d.connPool.WithContext(ctx).Where("currency_id = ? AND on_date = ?", curID, onDate).Find(&rates)

	if result.Error != nil {
		return nil, result.Error
	}

	return rates, nil
}

func (d DBClientStd) InsertRate(ctx context.Context, rate *schema.Rate) error {
	result := d.connPool.WithContext(ctx).Create(rate)

	return result.Error
}

func (d DBClientStd) InsertRates(ctx context.Context, rates []schema.Rate) error {
	result := d.connPool.WithContext(ctx).Create(rates)

	return result.Error
}
