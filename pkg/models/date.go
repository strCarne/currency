package models

import (
	"fmt"
	"time"
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

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}
