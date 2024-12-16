package model

// BS, PLの連結
func Consolidate(primaryBS, subsidiaryBS BS, primaryPL, subsidiaryPL PL) (BS, PL) {

	consolidatedBS := ConsolidateBS(primaryBS, subsidiaryBS)

	consolidatedPL := ConsolidatePL(primaryPL, subsidiaryPL)

	return consolidatedBS.applyPL(consolidatedPL), consolidatedPL
}

func ConsolidateBS(primaryBS, subsidiaryBS BS) BS {
	return BS{
		Debit: BSDebit{
			OtherAssets:     primaryBS.Debit.OtherAssets + subsidiaryBS.Debit.OtherAssets,
			SubsidiaryStock: 0,
			Goodwill:        primaryBS.Debit.SubsidiaryStock - subsidiaryBS.Credit.NetAssets.Sum(),
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: primaryBS.Credit.Liabilities.OtherLiabilities + subsidiaryBS.Credit.Liabilities.OtherLiabilities,
			},
			NetAssets: NetAssets{
				Capital:          primaryBS.Credit.NetAssets.Capital,
				CapitalSurplus:   primaryBS.Credit.NetAssets.CapitalSurplus,
				RetainedEarnings: primaryBS.Credit.NetAssets.RetainedEarnings,
			},
		},
	}
}

func ConsolidatePL(primaryPL, subsidiaryPL PL) PL {
	return PL{
		PLDebit{
			OtherExpenses: primaryPL.Debit.OtherExpenses + subsidiaryPL.Debit.OtherExpenses,
			NetIncome:     primaryPL.Debit.NetIncome + subsidiaryPL.Debit.NetIncome,
		},
		PLCredit{
			OtherIncome: primaryPL.Credit.OtherIncome + subsidiaryPL.Credit.OtherIncome,
		},
	}
}
