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

// 借方合計
func (debit *PLDebit) Sum() int {
	return debit.OtherExpenses + debit.NetIncome
}

// 貸方合計
func (credit *PLCredit) Sum() int {
	return credit.OtherIncome
}

// 貸借の一致を確認
func (bs *ProfitLoss) Validate() bool {
	debitTotal := bs.Debit.Sum()
	creditTotal := bs.Credit.Sum()
	return debitTotal == creditTotal
}
