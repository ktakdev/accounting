package main

import (
	"accounting/model"
	"encoding/json"
	"fmt"
)

func main() {
	primaryBS := model.BS{
		Debit: model.BSDebit{
			OtherAssets:     187000,
			SubsidiaryStock: 17000,
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
			OtherAssets: 48000,
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

	consolidatedBS, consolidatedPL := model.Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL, model.ConsolidateOptions{})

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

	if !isValidPL {
		fmt.Println("整合性のないPLです")
	}

	fmt.Println(string(jsonDataPL))
}
