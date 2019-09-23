package main

import (
	"fmt"
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraLTC(binance *binance.Binance) {
	BTC := 0.002

	BNBBTC, err := binance.GetOrderBook("BNBBTC")
	if err != nil {
		log.Println(err)
		return
	}

	LTCBNB, err := binance.GetOrderBook("LTCBNB")
	if err != nil {
		log.Println(err)
		return
	}

	LTCBTC, err := binance.GetOrderBook("LTCBTC")
	if err != nil {
		log.Println(err)
		return
	}

	BNB := orderCalc("buy", BTC, binance.StringToFloat(BNBBTC.Asks[0][0]))

	LTC := orderCalc("buy", BNB, binance.StringToFloat(LTCBNB.Asks[0][0]))

	LBTC := orderCalc("sell", LTC, binance.StringToFloat(LTCBTC.Bids[0][0]))

	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		fmt.Printf("LTC %.3f\n", percent)
	}
}
