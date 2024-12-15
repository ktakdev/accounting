package model

type Debit struct {
	OtherAssets     int `json:"諸資産"`
	SubsidiaryStock int `json:"子会社株式"`
}
type Credit struct {
	OtherLiabilities int `json:"諸負債"`
	Capital          int `json:"資本金"`
	CapitalSurplus   int `json:"資本剰余金"`
	RetainedEarnings int `json:"利益剰余金"`
}

type BalanceSheet struct {
	Debit  Debit  `json:"借方"`
	Credit Credit `json:"貸方"`
}

type BS = BalanceSheet
