package main

import (
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraBNB(binance *binance.Binance, coin string) {

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

	balances := GetBalances(binance)
	BTC := binance.StringToFloat(balances["BTC"])
	BNB, Bnb := orderCalc("buy", BTC, binance.StringToFloat(BNBBTC.Asks[0][0]))
	COIN, Coin := orderCalc("buy", BNB, binance.StringToFloat(COINBNB.Asks[0][0]))
	LBTC, Lbtc := orderCalc("sell", COIN, binance.StringToFloat(COINBTC.Bids[0][0]))
	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		log.Printf("%-10s - %-5s %.3f (%.8f - %.8f)\n", "Regra BNB", coin, percent, BTC, LBTC)
		binance.OrderMarket("BNBBTC", "BUY", Bnb)
		binance.OrderMarket(coin+"BNB", "BUY", Coin)
		binance.OrderMarket(coin+"BTC", "SELL", Lbtc)

	}
}
