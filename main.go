package main

import (
	"github.com/diogoagra/binance/binance"
)

func main() {
	binance := binance.New("zVpTDWFAdBthlDFalBlydV45qnp5kdaBxPLx1NKCVkNLyEGzHlshsMsPgjVBTxLz", "YnWoIZaxYlESbHSHTBc2l3BAc5Sp9GkYPZkeyCuG3VwjDXvjrfbqmMndeK9f4i4l", false)
	// binance.OrderTest("BTCUSDT", "BUY", 0.1, 10040.54)

	for {
		regraLTC(binance)
		regraXRP(binance)
	}
}
