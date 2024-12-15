package model

// BS, PLの連結
func Consolidate(primaryBS, subsidiaryBS BS, primaryPL, subsidiaryPL PL) (BS, PL) {
	consolidatedBS := BS{
		OtherAssets:      primaryBS.OtherAssets + subsidiaryBS.OtherAssets,
		SubsidiaryStock:  0,
		OtherLiabilities: primaryBS.OtherLiabilities + subsidiaryBS.OtherLiabilities,
		Capital:          primaryBS.Capital,
		CapitalSurplus:   primaryBS.CapitalSurplus,
		RetainedEarnings: primaryBS.RetainedEarnings + primaryPL.NetIncome + subsidiaryPL.NetIncome,
	}
	consolidatedPL := PL{
		NetIncome:     primaryPL.NetIncome + subsidiaryPL.NetIncome,
		OtherExpenses: primaryPL.OtherExpenses + subsidiaryPL.OtherExpenses,
		OtherIncome:   primaryPL.OtherIncome + subsidiaryPL.OtherIncome,
	}

	return consolidatedBS, consolidatedPL
}
