package main

import (
	"log"

	"github.com/diogoagra/binance/binance"
)

func regraUSDT(binance *binance.Binance, coin string) {

	balances := GetBalances(binance)

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

	BTC := binance.StringToFloat(balances["BTC"])
	USDT, Usdt := orderCalc("sell", BTC, binance.StringToFloat(BTCUSDT.Bids[0][0]))
	COIN, Coin := orderCalc("buy", USDT, binance.StringToFloat(COINUSDT.Asks[0][0]))
	LBTC, Lbtc := orderCalc("sell", COIN, binance.StringToFloat(COINBTC.Bids[0][0]))

	percent := (LBTC - BTC) * 100 / BTC

	if percent > 0 {
		log.Printf("%-10s - %-5s %.3f (%.8f - %.8f)\n", "Regra USDT", coin, percent, BTC, LBTC)
		binance.OrderMarket("BTCUSDT", "SELL", Usdt)
		binance.OrderMarket(coin+"USDT", "BUY", Coin)
		binance.OrderMarket(coin+"BTC", "SELL", Lbtc)
	}
}
