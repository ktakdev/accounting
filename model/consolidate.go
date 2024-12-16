package model

type ConsolidateOptions struct {
	ContorollingInterestRatio float64 `json:"親会社が保有している株式の割合"`
	SubsidiaryBSDiff          *BS     `json:"子会社のBS時価変動分"`
}

// BS, PLの連結
func Consolidate(primaryBS, subsidiaryBS BS, primaryPL, subsidiaryPL PL, opts ConsolidateOptions) (BS, PL) {
	consolidatedBS := ConsolidateBS(primaryBS, subsidiaryBS, opts)
	consolidatedPL := ConsolidatePL(primaryPL, subsidiaryPL, opts)

	return consolidatedBS.applyPL(consolidatedPL), consolidatedPL
}

func ConsolidateBS(primaryBS, subsidiaryBS BS, opts ConsolidateOptions) BS {
	// 親会社の保有割合の指定なし
	if opts.ContorollingInterestRatio == 0 {
		// すべて保有していると考える
		opts.ContorollingInterestRatio = 1
	}

	// 子会社BSと時価評価に差があるならそれを適用
	subCurrentValueBS := subsidiaryBS
	if opts.SubsidiaryBSDiff != nil {
		subCurrentValueBS = subsidiaryBS.Add(*opts.SubsidiaryBSDiff)
	}

	netAssetsSum := subCurrentValueBS.Credit.NetAssets.Sum()
	controllingInterests := netAssetsSum * opts.ContorollingInterestRatio
	nonControllingInterests := netAssetsSum - controllingInterests

	return BS{
		Debit: BSDebit{
			OtherAssets:     primaryBS.Debit.OtherAssets + subCurrentValueBS.Debit.OtherAssets,
			Land:            primaryBS.Debit.Land + subCurrentValueBS.Debit.Land,
			SubsidiaryStock: 0,
			Goodwill:        primaryBS.Debit.SubsidiaryStock - controllingInterests,
		},
		Credit: BSCredit{
			Liabilities: Liabilities{
				OtherLiabilities: primaryBS.Credit.Liabilities.OtherLiabilities + subCurrentValueBS.Credit.Liabilities.OtherLiabilities,
			},
			NetAssets: NetAssets{
				Capital:          primaryBS.Credit.NetAssets.Capital,
				CapitalSurplus:   primaryBS.Credit.NetAssets.CapitalSurplus,
				RetainedEarnings: primaryBS.Credit.NetAssets.RetainedEarnings,
				NCI:              nonControllingInterests,
			},
		},
	}
}

func ConsolidatePL(primaryPL, subsidiaryPL PL, opts ConsolidateOptions) PL {
	// 親会社の保有割合の指定なし
	if opts.ContorollingInterestRatio == 0 {
		// すべて保有していると考える
		opts.ContorollingInterestRatio = 1
	}

	CINetIncome := subsidiaryPL.Debit.NetIncome * opts.ContorollingInterestRatio
	NCINetIncome := subsidiaryPL.Debit.NetIncome - CINetIncome

	return PL{
		PLDebit{
			OtherExpenses: primaryPL.Debit.OtherExpenses + subsidiaryPL.Debit.OtherExpenses,
			NetIncome:     primaryPL.Debit.NetIncome + CINetIncome,
			NCINetIncome:  NCINetIncome,
		},
		PLCredit{
			OtherIncome: primaryPL.Credit.OtherIncome + subsidiaryPL.Credit.OtherIncome,
			NCIChange:   NCINetIncome,
		},
	}
}
