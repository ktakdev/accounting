package model

// BS, PLの連結
func Consolidate(primaryBS, subsidiaryBS BS, primaryPL, subsidiaryPL PL) (BS, PL) {

	consolidatedBS := BS{
		Debit: BSDebit{
			OtherAssets:     primaryBS.Debit.OtherAssets + subsidiaryBS.Debit.OtherAssets,
			SubsidiaryStock: 0,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: primaryBS.Credit.Liabilities.OtherLiabilities + subsidiaryBS.Credit.Liabilities.OtherLiabilities,
			},
			NetAssets: NetAssets{
				Capital:          primaryBS.Credit.NetAssets.Capital,
				CapitalSurplus:   primaryBS.Credit.NetAssets.CapitalSurplus,
				RetainedEarnings: primaryBS.Credit.NetAssets.RetainedEarnings + primaryPL.Debit.NetIncome + subsidiaryPL.Debit.NetIncome,
			},
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
