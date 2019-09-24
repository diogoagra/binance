package main

import (
	"log"

	"github.com/diogoagra/binance/binance"
)

// comprar eth com btc
// com eth comprar qualquer moeda
// com qualquer moeda vender por btc
// verificar saldo

func regraETH(binance *binance.Binance, coin string) {

	ETHBTC, err := binance.GetOrderBook("ETHBTC")
	if err != nil {
		log.Println(err)
		return
	}
	if len(ETHBTC.Asks) == 0 {
		return
	}

	COINETH, err := binance.GetOrderBook(coin + "ETH")
	if err != nil {
		log.Println(err)
		return
	}
	if len(COINETH.Asks) == 0 {
		return
	}

	COINBTC, err := binance.GetOrderBook(coin + "BTC")
	if err != nil {
		log.Println(err)
		return
	}
	if len(COINBTC.Bids) == 0 {
		return
	}

	BTC := 0.002
	ETH := orderCalc("buy", BTC, binance.StringToFloat(ETHBTC.Asks[0][0]))
	COIN := orderCalc("buy", ETH, binance.StringToFloat(COINETH.Asks[0][0]))
	LBTC := orderCalc("sell", COIN, binance.StringToFloat(COINBTC.Bids[0][0]))

	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		log.Printf("%-10s - %-5s %.3f (%.8f - %.8f)\n", "Regra ETH", coin, percent, BTC, LBTC)
	}
}
