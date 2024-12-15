package model

type BalanceSheet struct {
	SubsidiaryStock  int `json:"子会社株式"`
	OtherAssets      int `json:"諸資産"`
	Capital          int `json:"資本金"`
	CapitalSurplus   int `json:"資本剰余金"`
	RetainedEarnings int `json:"利益剰余金"`
	OtherLiabilities int `json:"諸負債"`
}

type BS = BalanceSheet
