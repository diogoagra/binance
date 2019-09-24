package main

import (
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraUSDT(binance *binance.Binance, coin string) {

	BTCUSDT, err := binance.GetOrderBook("BTCUSDT")
	if err != nil {
		log.Println(err)
		return
	}
	if len(BTCUSDT.Bids) == 0 {
		return
	}

	COINUSDT, err := binance.GetOrderBook(coin + "USDT")
	if err != nil {
		log.Println(err)
		return
	}
	if len(COINUSDT.Asks) == 0 {
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
	USDT := orderCalc("sell", BTC, binance.StringToFloat(BTCUSDT.Bids[0][0]))
	COIN := orderCalc("buy", USDT, binance.StringToFloat(COINUSDT.Asks[0][0]))
	LBTC := orderCalc("sell", COIN, binance.StringToFloat(COINBTC.Bids[0][0]))

	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		log.Printf("%-10s - %-5s %.3f (%.8f - %.8f)\n", "Regra USDT", coin, percent, BTC, LBTC)

		balances = GetBalances(binance)
		binance.OrderMarket("BTCUSDT", "SELL", balances["BTC"])
		balances = GetBalances(binance)
		binance.OrderMarket(coin+"USDT", "BUY", balances["USDT"])
		balances = GetBalances(binance)
		binance.OrderMarket(coin+"BTC", "SELL", balances[coin])

	}
}
