package model

type ProfitAndLossStatement struct {
	NetIncome     int `json:"当期純利益"`
	OtherExpenses int `json:"諸費用"`
	OtherIncome   int `json:"諸収益"`
}

type PL = ProfitAndLossStatement
