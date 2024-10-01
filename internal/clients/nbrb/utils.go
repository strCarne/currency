package nbrb

import (
	"net/url"
	"strconv"

	"github.com/strCarne/currency/pkg/models"
)

func evaluateQueryParams(
	onDate *models.Date,
	periodicity *Periodicity,
	paramMode *ParamMode,
) url.Values {
	queryParams := url.Values{}

	if onDate != nil {
		queryParams.Add("ondate", onDate.String())
	}

	if periodicity != nil {
		queryParams.Add("periodicity", strconv.Itoa(int(*periodicity)))
	}

	if paramMode != nil {
		queryParams.Add("mode", strconv.Itoa(int(*paramMode)))
	}

	return queryParams
}
