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
			Liabilities: Liabilities{
				OtherLiabilities: 86000,
			},
			NetAssets: NetAssets{
				Capital:          55000,
				CapitalSurplus:   10000,
				RetainedEarnings: 23000,
			},
		},
	}

	subsidiaryBS := BS{
		Debit: BSDebit{
			OtherAssets: 48000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 19000,
			},
			NetAssets: NetAssets{
				Capital:          10000,
				CapitalSurplus:   2000,
				RetainedEarnings: 5000,
			},
		},
	}

	primaryPL := PL{
		Debit: PLDebit{
			OtherExpenses: 570000,
			NetIncome:     30000,
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

	consolidatedBS, consolidatedPL := Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL, ConsolidateOptions{})

	// BSテスト

	if consolidatedBS.Debit.OtherAssets != 235000 {
		t.Errorf("諸資産 = %v,  want %v", consolidatedBS.Debit.OtherAssets, 235000)
	}

	if consolidatedBS.Debit.SubsidiaryStock != 0 {
		t.Errorf("子会社株式 = %v,  want %v", consolidatedBS.Debit.SubsidiaryStock, 0)
	}

	if consolidatedBS.Credit.Liabilities.OtherLiabilities != 105000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedBS.Credit.Liabilities.OtherLiabilities, 105000)
	}

	if consolidatedBS.Credit.NetAssets.Capital != 55000 {
		t.Errorf("資本金 = %v,  want %v", consolidatedBS.Credit.NetAssets.Capital, 55000)
	}

	if consolidatedBS.Credit.NetAssets.CapitalSurplus != 10000 {
		t.Errorf("資本剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.CapitalSurplus, 10000)
	}

	if consolidatedBS.Credit.NetAssets.RetainedEarnings != 65000 {
		t.Errorf("利益剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.RetainedEarnings, 65000)
	}

	if consolidatedBS.Credit.NetAssets.NCI != 0 {
		t.Errorf("被支配株主持分 = %v,  want %v", consolidatedBS.Credit.NetAssets.NCI, 0)
	}

	// PLテスト

	if consolidatedPL.Debit.OtherExpenses != 708000 {
		t.Errorf("諸費用 = %v,  want %v", consolidatedPL.Debit.OtherExpenses, 708000)
	}

	if consolidatedPL.Debit.NetIncome != 42000 {
		t.Errorf("当期純損益 = %v,  want %v", consolidatedPL.Debit.NetIncome, 42000)
	}

	if consolidatedPL.Credit.OtherIncome != 750000 {
		t.Errorf("諸収益 = %v,  want %v", consolidatedPL.Credit.OtherIncome, 750000)
	}
}

// 29-1-15
func TestConsolidateGoodwill(t *testing.T) {

	primaryBS := BS{
		Debit: BSDebit{
			OtherAssets:     154000,
			SubsidiaryStock: 20000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 86000,
			},
			NetAssets: NetAssets{
				Capital:          55000,
				CapitalSurplus:   10000,
				RetainedEarnings: 23000,
			},
		},
	}

	subsidiaryBS := BS{
		Debit: BSDebit{
			OtherAssets: 36000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 19000,
			},
			NetAssets: NetAssets{
				Capital:          10000,
				CapitalSurplus:   2000,
				RetainedEarnings: 5000,
			},
		},
	}

	consolidatedBS := ConsolidateBS(primaryBS, subsidiaryBS, ConsolidateOptions{})

	// BSテスト

	if consolidatedBS.Debit.OtherAssets != 190000 {
		t.Errorf("諸資産 = %v,  want %v", consolidatedBS.Debit.OtherAssets, 190000)
	}

	if consolidatedBS.Debit.SubsidiaryStock != 0 {
		t.Errorf("子会社株式 = %v,  want %v", consolidatedBS.Debit.SubsidiaryStock, 0)
	}

	if consolidatedBS.Debit.Goodwill != 3000 {
		t.Errorf("のれん = %v,  want %v", consolidatedBS.Debit.Goodwill, 3000)
	}

	if consolidatedBS.Credit.Liabilities.OtherLiabilities != 105000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedBS.Credit.Liabilities.OtherLiabilities, 105000)
	}

	if consolidatedBS.Credit.NetAssets.Capital != 55000 {
		t.Errorf("資本金 = %v,  want %v", consolidatedBS.Credit.NetAssets.Capital, 55000)
	}

	if consolidatedBS.Credit.NetAssets.CapitalSurplus != 10000 {
		t.Errorf("資本剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.CapitalSurplus, 10000)
	}

	if consolidatedBS.Credit.NetAssets.RetainedEarnings != 23000 {
		t.Errorf("利益剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.RetainedEarnings, 23000)
	}

	if consolidatedBS.Credit.NetAssets.NCI != 0 {
		t.Errorf("被支配株主持分 = %v,  want %v", consolidatedBS.Credit.NetAssets.NCI, 0)
	}
}

// 29-1-29
func TestConsolidateBSWithNCI(t *testing.T) {

	primaryBS := BS{
		Debit: BSDebit{
			OtherAssets:     108000,
			Land:            54000,
			SubsidiaryStock: 12000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 86000,
			},
			NetAssets: NetAssets{
				Capital:          55000,
				CapitalSurplus:   10000,
				RetainedEarnings: 23000,
			},
		},
	}

	subsidiaryBS := BS{
		Debit: BSDebit{
			OtherAssets: 30000,
			Land:        6000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 19000,
			},
			NetAssets: NetAssets{
				Capital:          10000,
				CapitalSurplus:   2000,
				RetainedEarnings: 5000,
			},
		},
	}

	// 連結条件
	opts := ConsolidateOptions{
		// 60%保有
		CIRatio: 0.6,
		// 子会社の土地が500円評価増
		SubsidiaryBSDiff: &BS{
			BSDebit{
				Land: 500,
			},
			BSCredit{
				NetAssets: NetAssets{
					FairValueDiff: 500,
				},
			},
		},
	}
	consolidatedBS := ConsolidateBS(primaryBS, subsidiaryBS, opts)

	// BSテスト

	if consolidatedBS.Debit.OtherAssets != 138000 {
		t.Errorf("諸資産 = %v,  want %v", consolidatedBS.Debit.OtherAssets, 138000)
	}

	if consolidatedBS.Debit.SubsidiaryStock != 0 {
		t.Errorf("子会社株式 = %v,  want %v", consolidatedBS.Debit.SubsidiaryStock, 0)
	}

	if consolidatedBS.Debit.Goodwill != 1500 {
		t.Errorf("のれん = %v,  want %v", consolidatedBS.Debit.Goodwill, 1500)
	}

	if consolidatedBS.Credit.Liabilities.OtherLiabilities != 105000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedBS.Credit.Liabilities.OtherLiabilities, 105000)
	}

	if consolidatedBS.Credit.NetAssets.Capital != 55000 {
		t.Errorf("資本金 = %v,  want %v", consolidatedBS.Credit.NetAssets.Capital, 55000)
	}

	if consolidatedBS.Credit.NetAssets.CapitalSurplus != 10000 {
		t.Errorf("資本剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.CapitalSurplus, 10000)
	}

	if consolidatedBS.Credit.NetAssets.NCI != 7000 {
		t.Errorf("被支配株主持分 = %v,  want %v", consolidatedBS.Credit.NetAssets.NCI, 7000)
	}
}

// 29-1-33
func TestConsolidateWithNCI(t *testing.T) {

	primaryBS := BS{
		Debit: BSDebit{
			OtherAssets:     138000,
			Land:            54000,
			SubsidiaryStock: 12000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 86000,
			},
			NetAssets: NetAssets{
				Capital:          55000,
				CapitalSurplus:   10000,
				RetainedEarnings: 23000,
			},
		},
	}

	subsidiaryBS := BS{
		Debit: BSDebit{
			OtherAssets: 42000,
			Land:        6000,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: 19000,
			},
			NetAssets: NetAssets{
				Capital:          10000,
				CapitalSurplus:   2000,
				RetainedEarnings: 5000,
			},
		},
	}

	primaryPL := PL{
		Debit: PLDebit{
			OtherExpenses: 570000,
			NetIncome:     30000,
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

	// 連結条件
	opts := ConsolidateOptions{
		// 60%保有
		CIRatio: 0.6,
		// 子会社の土地が500円評価増
		SubsidiaryBSDiff: &BS{
			BSDebit{
				Land: 500,
			},
			BSCredit{
				NetAssets: NetAssets{
					FairValueDiff: 500,
				},
			},
		},
	}
	consolidatedBS, consolidatedPL := Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL, opts)

	// BSテスト

	if consolidatedBS.Debit.OtherAssets != 180000 {
		t.Errorf("諸資産 = %v,  want %v", consolidatedBS.Debit.OtherAssets, 180000)
	}

	if consolidatedBS.Debit.SubsidiaryStock != 0 {
		t.Errorf("子会社株式 = %v,  want %v", consolidatedBS.Debit.SubsidiaryStock, 0)
	}

	if consolidatedBS.Debit.Goodwill != 1425 {
		t.Errorf("のれん = %v,  want %v", consolidatedBS.Debit.Goodwill, 1425)
	}

	if consolidatedBS.Credit.Liabilities.OtherLiabilities != 105000 {
		t.Errorf("諸負債 = %v,  want %v", consolidatedBS.Credit.Liabilities.OtherLiabilities, 105000)
	}

	if consolidatedBS.Credit.NetAssets.Capital != 55000 {
		t.Errorf("資本金 = %v,  want %v", consolidatedBS.Credit.NetAssets.Capital, 55000)
	}

	if consolidatedBS.Credit.NetAssets.CapitalSurplus != 10000 {
		t.Errorf("資本剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.CapitalSurplus, 10000)
	}

	if consolidatedBS.Credit.NetAssets.RetainedEarnings != 60125 {
		t.Errorf("利益剰余金 = %v,  want %v", consolidatedBS.Credit.NetAssets.RetainedEarnings, 60125)
	}

	if consolidatedBS.Credit.NetAssets.NCI != 11800 {
		t.Errorf("被支配株主持分 = %v,  want %v", consolidatedBS.Credit.NetAssets.NCI, 11800)
	}

	// PLテスト

	if consolidatedPL.Debit.OtherExpenses != 708000 {
		t.Errorf("諸費用 = %v,  want %v", consolidatedPL.Debit.OtherExpenses, 708000)
	}

	if consolidatedPL.Debit.GoodwillAmortization != 75 {
		t.Errorf("のれん償却 = %v,  want %v", consolidatedPL.Debit.GoodwillAmortization, 75)
	}

	if consolidatedPL.Debit.NetIncome != 37125 {
		t.Errorf("当期純損益 = %v,  want %v", consolidatedPL.Debit.NetIncome, 37125)
	}

	if consolidatedPL.Debit.NCINetIncome != 4800 {
		t.Errorf("被支配株主に帰属する当期純損益 = %v,  want %v", consolidatedPL.Debit.NCINetIncome, 4800)
	}

	if consolidatedPL.Credit.OtherIncome != 750000 {
		t.Errorf("諸収益 = %v,  want %v", consolidatedPL.Credit.OtherIncome, 750000)
	}
}
