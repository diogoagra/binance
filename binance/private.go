package binance

import (
	"fmt"
	"strconv"
	"time"
)

// OrderTest -
func (b *Binance) OrderTest(symbol, side string, quantity, price float64) {
	timestamp := strconv.Itoa(int(unixMillis(time.Now())))
	parameters := fmt.Sprintf("symbol=%s&side=%s&type=LIMIT&timeInForce=GTC&quantity=%.8f&price=%.8f&recvWindow=5000&timestamp=%s", symbol, side, quantity, price, timestamp)
	response, err := b.GetURL("/api/v3/order/test?"+fmt.Sprintf("%s&signature=%s", parameters, b.Sign(parameters)), "POST", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(response))
}
