package model

import (
	"testing"
)

func TestConsolidate(t *testing.T) {
	primaryBS := BS{
		Debit: BSDebit{
			OtherAssets:     187000,
			SubsidiaryStock: 17000,
		},
		Credit: BSCredit{
			OtherLiabilities: 86000,
			Capital:          55000,
			CapitalSurplus:   10000,
			RetainedEarnings: 23000,
		},
	}

	subsidiaryBS := BS{
		Debit: BSDebit{
			OtherAssets: 48000,
		},
		Credit: BSCredit{
			OtherLiabilities: 19000,
			Capital:          10000,
			CapitalSurplus:   2000,
			RetainedEarnings: 5000,
		},
	}

	primaryPL := PL{
		Debit: PLDebit{OtherExpenses: 570000,
			NetIncome: 30000,
		},
		Credit: PLCredit{
			OtherIncome: 600000,
		},
	}

	subsidiaryPL := PL{
		Debit: PLDebit{
			OtherExpenses: 138000,
			NetIncome:     12000,
		},
		Credit: PLCredit{
			OtherIncome: 150000,
		},
	}

	consolidatedBS, consolidatedPL := Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL)

	// BSテスト

	if consolidatedBS.Debit.OtherAssets != 235000 {
		t.Errorf("諸資産 = %v,  want %v", consolidatedBS.Debit.OtherAssets, 235000)
	}

	if consolidatedBS.Debit.SubsidiaryStock != 0 {
		t.Errorf("子会社株式 = %v,  want %v", consolidatedBS.Debit.SubsidiaryStock, 0)
	}

	if consolidatedBS.Credit.OtherLiabilities != 105000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedBS.Credit.OtherLiabilities, 105000)
	}

	if consolidatedBS.Credit.Capital != 55000 {
		t.Errorf("資本金 = %v,  want %v", consolidatedBS.Credit.Capital, 55000)
	}

	if consolidatedBS.Credit.CapitalSurplus != 10000 {
		t.Errorf("資本剰余金 = %v,  want %v", consolidatedBS.Credit.CapitalSurplus, 10000)
	}

	if consolidatedBS.Credit.RetainedEarnings != 65000 {
		t.Errorf("利益剰余金 = %v,  want %v", consolidatedBS.Credit.RetainedEarnings, 65000)
	}

	// PLテスト

	if consolidatedPL.Debit.OtherExpenses != 708000 {
		t.Errorf("諸費用 = %v,  want %v", consolidatedPL.Debit.OtherExpenses, 708000)
	}

	if consolidatedPL.Debit.NetIncome != 42000 {
		t.Errorf("当期純利益 = %v,  want %v", consolidatedPL.Debit.OtherExpenses, 42000)
	}

	if consolidatedPL.Credit.OtherIncome != 750000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedPL.Credit.OtherIncome, 750000)
	}
}
