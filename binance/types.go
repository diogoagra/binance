package binance

// OrderBookStruct -
type OrderBookStruct struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
