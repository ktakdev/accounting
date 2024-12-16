package main

import (
	"accounting/model"
	"encoding/json"
	"fmt"
)

func main() {
	primaryBS := model.BS{
		Debit: model.BSDebit{
			OtherAssets:     138000,
			Land:            54000,
			SubsidiaryStock: 12000,
		},
		Credit: model.BSCredit{
			Liabilities: model.Liabilities{
				OtherLiabilities: 86000,
			},
			NetAssets: model.NetAssets{
				Capital:          55000,
				CapitalSurplus:   10000,
				RetainedEarnings: 23000,
			},
		},
	}

	subsidiaryBS := model.BS{
		Debit: model.BSDebit{
			OtherAssets: 42000,
			Land:        6000,
		},
		Credit: model.BSCredit{
			Liabilities: model.Liabilities{
				OtherLiabilities: 19000,
			},
			NetAssets: model.NetAssets{
				Capital:          10000,
				CapitalSurplus:   2000,
				RetainedEarnings: 5000,
			},
		},
	}

	primaryPL := model.PL{
		Debit: model.PLDebit{
			OtherExpenses: 570000,
			NetIncome:     30000,
		},
		Credit: model.PLCredit{
			OtherIncome: 600000,
		},
	}

	subsidiaryPL := model.PL{
		Debit: model.PLDebit{
			OtherExpenses: 138000,
			NetIncome:     12000,
		},
		Credit: model.PLCredit{
			OtherIncome: 150000,
		},
	}

	// 連結条件
	opts := model.ConsolidateOptions{
		// 60%保有
		CIRatio: 0.6,
		// 子会社の土地が500円評価増
		SubsidiaryBSDiff: &model.BS{
			Debit: model.BSDebit{
				Land: 500,
			},
			Credit: model.BSCredit{
				NetAssets: model.NetAssets{
					FairValueDiff: 500,
				},
			},
		},
	}
	consolidatedBS, consolidatedPL := model.Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL, opts)

	jsonDataBS, err := json.Marshal(consolidatedBS)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonDataBS))
	isValidBS := consolidatedBS.Validate()

	if !isValidBS {
		fmt.Println("整合性のないBSです")
	}

	jsonDataPL, err := json.Marshal(consolidatedPL)
	isValidPL := consolidatedPL.Validate()
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonDataPL))
	if !isValidPL {
		fmt.Println("整合性のないPLです")
	}
}
