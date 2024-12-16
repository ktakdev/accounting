package model

type PLDebit struct {
	OtherExpenses        float64 `json:"諸費用,omitempty"`
	NetIncome            float64 `json:"当期純損益,omitempty"`
	NCINetIncome         float64 `json:"被支配株主に帰属する当期純損益,omitempty"`
	GoodwillAmortization float64 `json:"のれん償却,omitempty"`
}
type PLCredit struct {
	OtherIncome float64 `json:"諸収益,omitempty"`
	NCIChange   float64 `json:"被支配株主持分-当期変動額,omitempty"`
}

type ProfitLoss struct {
	Debit  PLDebit  `json:"借方,omitempty"`
	Credit PLCredit `json:"貸方,omitempty"`
}

type PL = ProfitLoss

// 借方合計
func (debit *PLDebit) Sum() float64 {
	return debit.OtherExpenses + debit.NetIncome + debit.NCINetIncome
}

// 貸方合計
func (credit *PLCredit) Sum() float64 {
	return credit.OtherIncome
}

// 貸借の一致を確認
func (pl *ProfitLoss) Validate() bool {
	debitTotal := pl.Debit.Sum()
	creditTotal := pl.Credit.Sum()
	return debitTotal == creditTotal
}
