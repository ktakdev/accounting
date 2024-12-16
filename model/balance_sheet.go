package model

type BalanceSheet struct {
	Debit  BSDebit  `json:"借方"`
	Credit BSCredit `json:"貸方"`
}
type BSDebit struct {
	OtherAssets     float64 `json:"諸資産"`
	Land            float64 `json:"土地"`
	SubsidiaryStock float64 `json:"子会社株式"`
	Goodwill        float64 `json:"のれん"`
}

type BSCredit struct {
	Liabilities Liabilities `json:"負債"`
	NetAssets   NetAssets   `json:"純資産"`
}

type Liabilities struct {
	OtherLiabilities float64 `json:"諸負債"`
}

type NetAssets struct {
	Capital          float64 `json:"資本金"`
	CapitalSurplus   float64 `json:"資本剰余金"`
	RetainedEarnings float64 `json:"利益剰余金"`
	NCI              float64 `json:"被支配株主持分"`
	FairValueDiff    float64 `json:"評価差額"`
}

type BS = BalanceSheet

// 借方合計
func (debit *BSDebit) Sum() float64 {
	return debit.OtherAssets + debit.Land + debit.SubsidiaryStock
}

// 借方の足し算
func (debit *BSDebit) Add(debit2 BSDebit) BSDebit {
	return BSDebit{
		OtherAssets:     debit.OtherAssets + debit2.OtherAssets,
		Land:            debit.Land + debit2.Land,
		SubsidiaryStock: debit.SubsidiaryStock + debit2.SubsidiaryStock,
		Goodwill:        debit.Goodwill + debit2.Goodwill,
	}
}

// 貸方合計
func (credit *BSCredit) Sum() float64 {
	return credit.Liabilities.Sum() + credit.NetAssets.Sum()
}

// 貸方の足し算
func (credit *BSCredit) Add(credit2 BSCredit) BSCredit {
	return BSCredit{
		Liabilities: credit.Liabilities.Add(credit2.Liabilities),
		NetAssets:   credit.NetAssets.Add(credit2.NetAssets),
	}
}

// 負債合計
func (liabilies *Liabilities) Sum() float64 {
	return liabilies.OtherLiabilities
}

// 負債の足し算
func (liabilies *Liabilities) Add(liabilies2 Liabilities) Liabilities {
	return Liabilities{
		OtherLiabilities: liabilies.OtherLiabilities + liabilies2.OtherLiabilities,
	}
}

// 純資産合計
func (netAssets *NetAssets) Sum() float64 {
	return netAssets.Capital + netAssets.CapitalSurplus + netAssets.RetainedEarnings
}

// 純資産の足し算
func (netAssets *NetAssets) Add(netAssets2 NetAssets) NetAssets {
	return NetAssets{
		Capital:          netAssets.Capital + netAssets2.Capital,
		CapitalSurplus:   netAssets.CapitalSurplus + netAssets2.CapitalSurplus,
		RetainedEarnings: netAssets.RetainedEarnings + netAssets2.RetainedEarnings,
		NCI:              netAssets.NCI + netAssets2.NCI,
		FairValueDiff:    netAssets.FairValueDiff + netAssets2.FairValueDiff,
	}
}

// BSの足し算
func (bs *BalanceSheet) Add(bs2 BalanceSheet) BalanceSheet {
	return BS{
		Debit:  bs.Debit.Add(bs2.Debit),
		Credit: bs.Credit.Add(bs2.Credit),
	}
}

// 貸借の一致を確認
func (bs *BalanceSheet) Validate() bool {
	debitTotal := bs.Debit.Sum()
	creditTotal := bs.Credit.Sum()
	return debitTotal == creditTotal
}

// PLを適用する
func (bs *BalanceSheet) applyPL(pl PL) BalanceSheet {
	return BalanceSheet{
		Debit: bs.Debit,
		Credit: BSCredit{
			Liabilities: bs.Credit.Liabilities,
			NetAssets: NetAssets{
				Capital:          bs.Credit.NetAssets.Capital,
				CapitalSurplus:   bs.Credit.NetAssets.CapitalSurplus,
				RetainedEarnings: bs.Credit.NetAssets.RetainedEarnings + pl.Debit.NetIncome,
			},
		},
	}
}
