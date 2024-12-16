package model

type BalanceSheet struct {
	Debit  BSDebit  `json:"借方"`
	Credit BSCredit `json:"貸方"`
}
type BSDebit struct {
	OtherAssets     int `json:"諸資産"`
	SubsidiaryStock int `json:"子会社株式"`
	Goodwill        int `json:"のれん"`
}

type BSCredit struct {
	Liabilities Liabilities `json:"負債"`
	NetAssets   NetAssets   `json:"純資産"`
}

type Liabilities struct {
	OtherLiabilities int `json:"諸負債"`
}

type NetAssets struct {
	Capital          int `json:"資本金"`
	CapitalSurplus   int `json:"資本剰余金"`
	RetainedEarnings int `json:"利益剰余金"`
}

type BS = BalanceSheet

// 借方合計
func (debit *BSDebit) Sum() int {
	return debit.OtherAssets + debit.SubsidiaryStock
}

// 貸方合計
func (credit *BSCredit) Sum() int {
	return credit.Liabilities.Sum() + credit.NetAssets.Sum()
}

// 負債合計
func (liabilies *Liabilities) Sum() int {
	return liabilies.OtherLiabilities
}

// 純資産合計
func (netAssets *NetAssets) Sum() int {
	return netAssets.Capital + netAssets.CapitalSurplus + netAssets.RetainedEarnings
}

// 貸借の一致を確認
func (bs *BalanceSheet) Validate() bool {
	debitTotal := bs.Debit.Sum()
	creditTotal := bs.Credit.Sum()
	return debitTotal == creditTotal
}
