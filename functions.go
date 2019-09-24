package main

import "github.com/diogoagra/binance/binance"

// GetBalances -
func GetBalances(binance *binance.Binance) (allBalances map[string]string) {

	allBalances = make(map[string]string)

	balances, err := binance.GetBalances()
	if err != nil {
		return
	}
	for _, b := range balances.Balances {
		allBalances[b.Asset] = b.Free
	}
	return
}
