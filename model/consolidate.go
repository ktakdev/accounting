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
			RetainedEarnings: primaryBS.Credit.RetainedEarnings + primaryPL.Debit.NetIncome + subsidiaryPL.Debit.NetIncome,
		},
	}

	consolidatedPL := PL{
		PLDebit{
			OtherExpenses: primaryPL.Debit.OtherExpenses + subsidiaryPL.Debit.OtherExpenses,
			NetIncome:     primaryPL.Debit.NetIncome + subsidiaryPL.Debit.NetIncome,
		},
		PLCredit{
			OtherIncome: primaryPL.Credit.OtherIncome + subsidiaryPL.Credit.OtherIncome,
		},
	}

	return consolidatedBS, consolidatedPL
}
