package model

type BalanceSheet struct {
	Capital int `json:"capital"`
}

// BSの連結
func Consolidate(primary, subsidiary BalanceSheet) BalanceSheet {
	return BalanceSheet{
		Capital: primary.Capital + subsidiary.Capital,
	}
}
