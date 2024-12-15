package model

// BS, PLの連結
func Consolidate(primaryBS, subsidiaryBS BS, primaryPL, subsidiaryPL PL) (BS, PL) {
	consolidatedBS := BS{
		Debit: BSDebit{
			OtherAssets:     primaryBS.Debit.OtherAssets + subsidiaryBS.Debit.OtherAssets,
			SubsidiaryStock: 0,
		},
		Credit: BSCredit{
			OtherLiabilities: primaryBS.Credit.OtherLiabilities + subsidiaryBS.Credit.OtherLiabilities,
			Capital:          primaryBS.Credit.Capital,
			CapitalSurplus:   primaryBS.Credit.CapitalSurplus,
			RetainedEarnings: primaryBS.Credit.RetainedEarnings + primaryPL.NetIncome + subsidiaryPL.NetIncome,
		},
	}
	consolidatedPL := PL{
		NetIncome:     primaryPL.NetIncome + subsidiaryPL.NetIncome,
		OtherExpenses: primaryPL.OtherExpenses + subsidiaryPL.OtherExpenses,
		OtherIncome:   primaryPL.OtherIncome + subsidiaryPL.OtherIncome,
	}

	return consolidatedBS, consolidatedPL
}
