package schema

//nolint:tagliatelle
type Rate struct {
	ID           int     `json:"Cur_ID"`
	Date         string  `json:"Date"`
	Abbreviation string  `json:"Cur_Abbreviation"`
	Scale        int     `json:"Cur_Scale"`
	Name         string  `json:"Cur_Name"`
	OfficialRate float64 `json:"Cur_OfficialRate"`
}
