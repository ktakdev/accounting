package model

// BSの連結
func Consolidate(primary, subsidiary BalanceSheet) BalanceSheet {
	return BalanceSheet{
		SubsidiaryStock:  0,
		Capital:          primary.Capital + subsidiary.Capital,
		CapitalSurplus:   primary.CapitalSurplus,
		RetainedEarnings: primary.RetainedEarnings,
	}
}
