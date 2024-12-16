package model

type PLDebit struct {
	OtherExpenses float64 `json:"諸費用,omitempty"`
	NetIncome     float64 `json:"当期純利益,omitempty"`
}
type PLCredit struct {
	OtherIncome float64 `json:"諸収益,omitempty"`
}

type ProfitLoss struct {
	Debit  PLDebit  `json:"借方,omitempty"`
	Credit PLCredit `json:"貸方,omitempty"`
}

type PL = ProfitLoss

// 借方合計
func (debit *PLDebit) Sum() float64 {
	return debit.OtherExpenses + debit.NetIncome
}

// 貸方合計
func (credit *PLCredit) Sum() float64 {
	return credit.OtherIncome
}

// 貸借の一致を確認
func (bs *ProfitLoss) Validate() bool {
	debitTotal := bs.Debit.Sum()
	creditTotal := bs.Credit.Sum()
	return debitTotal == creditTotal
}
