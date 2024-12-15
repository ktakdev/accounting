package main

import (
	"accounting/model"
	"encoding/json"
	"fmt"
)

func main() {
	bs := model.BalanceSheet{Capital: 100}

	jsonData, err := json.Marshal(bs)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
