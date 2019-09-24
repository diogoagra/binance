package main

import (
	"log"
	"reflect"

	"github.com/diogoagra/binance/binance"
)

func main() {
	binance := binance.New("zVpTDWFAdBthlDFalBlydV45qnp5kdaBxPLx1NKCVkNLyEGzHlshsMsPgjVBTxLz", "YnWoIZaxYlESbHSHTBc2l3BAc5Sp9GkYPZkeyCuG3VwjDXvjrfbqmMndeK9f4i4l", false)

	markets := make(map[string][]string)
	qtdaMin := make(map[string]string)

	marketsInfo, err := binance.ExchangeInfo()
	if err != nil {
		log.Println(err)
		return
	}

	for _, b := range marketsInfo.Symbols {

		qtdaMin[b.BaseAsset] = b.Filters[0].MinQty

		markets[b.BaseAsset] = append(markets[b.BaseAsset], b.QuoteAsset)
	}

	for {
		for a, b := range markets {
			if inArray("BTC", b) && inArray("ETH", b) {
				regraETH(binance, a)
			}
			if inArray("BTC", b) && inArray("BNB", b) {
				regraBNB(binance, a)
			}
			if inArray("BTC", b) && inArray("USDT", b) && inArray("BNB", b) {
				regraUSDT(binance, a)
			}
		}
	}
}

func inArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}
