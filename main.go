package main

import (
	"github.com/diogoagra/binance/binance"
)

func main() {
	binance := binance.New("zVpTDWFAdBthlDFalBlydV45qnp5kdaBxPLx1NKCVkNLyEGzHlshsMsPgjVBTxLz", "YnWoIZaxYlESbHSHTBc2l3BAc5Sp9GkYPZkeyCuG3VwjDXvjrfbqmMndeK9f4i4l", false)
	// binance.OrderTest("BTCUSDT", "BUY", 0.1, 10040.54)

	for {
		regra(binance, "LTC")
		regra(binance, "BAND")
		regra(binance, "EOS")
		regra(binance, "XRP")
		regra(binance, "ADA")
		regra(binance, "TRX")
		regra(binance, "WAVES")

	}
}
