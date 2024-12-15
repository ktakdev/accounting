package model

type BSDebit struct {
	OtherAssets     int `json:"諸資産"`
	SubsidiaryStock int `json:"子会社株式"`
}

type BSCredit struct {
	OtherLiabilities int `json:"諸負債"`
	Capital          int `json:"資本金"`
	CapitalSurplus   int `json:"資本剰余金"`
	RetainedEarnings int `json:"利益剰余金"`
}

type BalanceSheet struct {
	Debit  BSDebit  `json:"借方"`
	Credit BSCredit `json:"貸方"`
}

type BS = BalanceSheet

// 借方合計
func (debit *BSDebit) Sum() int {
	return debit.OtherAssets + debit.SubsidiaryStock
}

// 貸方合計
func (credit *BSCredit) Sum() int {
	return credit.OtherLiabilities + credit.Capital + credit.CapitalSurplus + credit.RetainedEarnings
}

// 貸借の一致を確認
func (bs *BalanceSheet) Validate() bool {
	debitTotal := bs.Debit.Sum()
	creditTotal := bs.Credit.Sum()
	return debitTotal == creditTotal
}
