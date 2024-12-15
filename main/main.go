package main

import (
	"accounting/model"
	"encoding/json"
	"fmt"
)

func main() {
	primary := model.BalanceSheet{SubsidiaryStock: 17000, Capital: 55000, CapitalSurplus: 10000, RetainedEarnings: 23000}
	secondary := model.BalanceSheet{Capital: 10000, CapitalSurplus: 2000, RetainedEarnings: 5000}

	consolidated := model.Consolidate(primary, secondary)

	jsonData, err := json.Marshal(consolidated)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
