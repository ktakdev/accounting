package model

type BalanceSheet struct {
	SubsidiaryStock  int `json:"子会社株式"`
	Capital          int `json:"資本金"`
	CapitalSurplus   int `json:"資本剰余金"`
	RetainedEarnings int `json:"利益剰余金"`
}

type BS = BalanceSheet
