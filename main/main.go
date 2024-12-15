package main

import (
	"accounting/model"
	"encoding/json"
	"fmt"
)

func main() {
	primaryBS := model.BS{
		Debit: model.Debit{
			OtherAssets:     187000,
			SubsidiaryStock: 17000,
		},
		Credit: model.Credit{
			OtherLiabilities: 86000,
			Capital:          55000,
			CapitalSurplus:   10000,
			RetainedEarnings: 23000,
		},
	}

	subsidiaryBS := model.BS{
		Debit: model.Debit{
			OtherAssets: 48000,
		},
		Credit: model.Credit{
			OtherLiabilities: 19000,
			Capital:          10000,
			CapitalSurplus:   2000,
			RetainedEarnings: 5000,
		},
	}

	primaryPL := model.PL{NetIncome: 30000, OtherExpenses: 570000, OtherIncome: 600000}
	subsidiaryPL := model.PL{NetIncome: 12000, OtherExpenses: 138000, OtherIncome: 150000}

	consolidatedBS, consolidatedPL := model.Consolidate(primaryBS, subsidiaryBS, primaryPL, subsidiaryPL)

	jsonDataBS, err := json.Marshal(consolidatedBS)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonDataBS))

	jsonDataPL, err := json.Marshal(consolidatedPL)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonDataPL))
}
