package model

type ProfitLoss struct {
	OtherExpenses int `json:"諸費用"`
	NetIncome     int `json:"当期純利益"`
	OtherIncome   int `json:"諸収益"`
}

type PL = ProfitLoss
