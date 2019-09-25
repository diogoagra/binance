package binance

import (
	"encoding/json"
	"fmt"
	"log"
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

// OrderMarket -
func (b *Binance) OrderMarket(symbol, side string, quantity float64) {
	timestamp := strconv.Itoa(int(unixMillis(time.Now())))
	parameters := fmt.Sprintf("symbol=%s&side=%s&type=MARKET&quantity=%.6f&recvWindow=5000&timestamp=%s", symbol, side, quantity, timestamp)
	response, err := b.GetURL("/api/v3/order?"+fmt.Sprintf("%s&signature=%s", parameters, b.Sign(parameters)), "POST", true)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(response))
}

// GetBalances -
func (b *Binance) GetBalances() (result GetBalancesStruct, err error) {
	timestamp := strconv.Itoa(int(unixMillis(time.Now())))
	parameters := fmt.Sprintf("recvWindow=5000&timestamp=%s", timestamp)
	response, err := b.GetURL("/api/v3/account?"+fmt.Sprintf("%s&signature=%s", parameters, b.Sign(parameters)), "GET", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
