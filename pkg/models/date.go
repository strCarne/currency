package models

import (
	"fmt"
	"time"

	"github.com/strCarne/currency/pkg/wrapper"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func DateNow() Date {
	year, month, day := time.Now().Date()

	return Date{
		Year:  year,
		Month: int(month),
		Day:   day,
	}
}

func DateParse(s string) (*Date, error) {
	parsedTime, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, wrapper.Wrap("models.DateParse", "failed to parse date", err)
	}

	return &Date{
		Year:  parsedTime.Year(),
		Month: int(parsedTime.Month()),
		Day:   parsedTime.Day(),
	}, nil
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02dT00:00:00", d.Year, d.Month, d.Day)
}
