package main

import (
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraBNB(binance *binance.Binance, coin string) {
	BTC := 0.002

	BNBBTC, err := binance.GetOrderBook("BNBBTC")
	if err != nil {
		log.Println(err)
		return
	}
	if len(BNBBTC.Asks) == 0 {
		return
	}

	COINBNB, err := binance.GetOrderBook(coin + "BNB")
	if err != nil {
		log.Println(err)
		return
	}
	if len(COINBNB.Asks) == 0 {
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

	BNB := orderCalc("buy", BTC, binance.StringToFloat(BNBBTC.Asks[0][0]))
	COIN := orderCalc("buy", BNB, binance.StringToFloat(COINBNB.Asks[0][0]))
	LBTC := orderCalc("sell", COIN, binance.StringToFloat(COINBTC.Bids[0][0]))
	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		log.Printf("%-10s - %-5s %.3f (%.8f - %.8f)\n", "Regra BNB", coin, percent, BTC, LBTC)
	}
}
