package binance

import (
	"encoding/json"
	"fmt"
)

// GetOrderBook -
func (b *Binance) GetOrderBook(symbol string) (result OrderBookStruct, err error) {
	response, err := b.GetURL(fmt.Sprintf("/api/v1/depth?symbol=%s", symbol), "GET", false)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
