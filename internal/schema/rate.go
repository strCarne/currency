package schema

import (
	"gorm.io/gorm"
)

//nolint:tagliatelle
type Rate struct {
	gorm.Model   `json:"-"`
	CurID        int     `gorm:"not null;uniqueIndex:idx_unique_curid_date"                  json:"Cur_ID"`
	Date         string  `gorm:"type:varchar(20);not null;uniqueIndex:idx_unique_curid_date" json:"Date"`
	Abbreviation string  `gorm:"type:varchar(5);not null"                                    json:"Cur_Abbreviation"`
	Scale        int     `gorm:"not null"                                                    json:"Cur_Scale"`
	Name         string  `gorm:"type:varchar(64);not null"                                   json:"Cur_Name"`
	OfficialRate float64 `gorm:"not null"                                                    json:"Cur_OfficialRate"`
}
