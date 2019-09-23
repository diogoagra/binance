package main

import (
	"fmt"
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraXRP(binance *binance.Binance) {
	BTC := 0.002

	BNBBTC, err := binance.GetOrderBook("BNBBTC")
	if err != nil {
		log.Println(err)
		return
	}

	XRPBNB, err := binance.GetOrderBook("XRPBNB")
	if err != nil {
		log.Println(err)
		return
	}

	XRPBTC, err := binance.GetOrderBook("XRPBTC")
	if err != nil {
		log.Println(err)
		return
	}

	BNB := orderCalc("buy", BTC, binance.StringToFloat(BNBBTC.Asks[0][0]))

	XRP := orderCalc("buy", BNB, binance.StringToFloat(XRPBNB.Asks[0][0]))

	LBTC := orderCalc("sell", XRP, binance.StringToFloat(XRPBTC.Bids[0][0]))

	percent := (LBTC - BTC) * 100 / BTC
	if percent > 0 {
		fmt.Printf("XRP %.3f\n", percent)
	}
}
