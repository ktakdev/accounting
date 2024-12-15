package model

type PLDebit struct {
	OtherExpenses int `json:"諸費用"`
	NetIncome     int `json:"当期純利益"`
}
type PLCredit struct {
	OtherIncome int `json:"諸収益"`
}

type ProfitLoss struct {
	Debit  PLDebit  `json:"借方"`
	Credit PLCredit `json:"貸方"`
}

type PL = ProfitLoss
