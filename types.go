package binance

// OrderBookStruct -
type OrderBookStruct struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// ExchangeInfoStruct -
type ExchangeInfoStruct struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                 string   `json:"symbol"`
		Status                 string   `json:"status"`
		BaseAsset              string   `json:"baseAsset"`
		BaseAssetPrecision     int      `json:"baseAssetPrecision"`
		QuoteAsset             string   `json:"quoteAsset"`
		QuotePrecision         int      `json:"quotePrecision"`
		OrderTypes             []string `json:"orderTypes"`
		IcebergAllowed         bool     `json:"icebergAllowed"`
		OcoAllowed             bool     `json:"ocoAllowed"`
		IsSpotTradingAllowed   bool     `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed bool     `json:"isMarginTradingAllowed"`
		Filters                []struct {
			FilterType       string `json:"filterType"`
			MinPrice         string `json:"minPrice,omitempty"`
			MaxPrice         string `json:"maxPrice,omitempty"`
			TickSize         string `json:"tickSize,omitempty"`
			MultiplierUp     string `json:"multiplierUp,omitempty"`
			MultiplierDown   string `json:"multiplierDown,omitempty"`
			AvgPriceMins     int    `json:"avgPriceMins,omitempty"`
			MinQty           string `json:"minQty,omitempty"`
			MaxQty           string `json:"maxQty,omitempty"`
			StepSize         string `json:"stepSize,omitempty"`
			MinNotional      string `json:"minNotional,omitempty"`
			ApplyToMarket    bool   `json:"applyToMarket,omitempty"`
			Limit            int    `json:"limit,omitempty"`
			MaxNumAlgoOrders int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
	} `json:"symbols"`
}

// GetBalancesStruct -
type GetBalancesStruct struct {
	MakerCommission  int  `json:"makerCommission"`
	TakerCommission  int  `json:"takerCommission"`
	BuyerCommission  int  `json:"buyerCommission"`
	SellerCommission int  `json:"sellerCommission"`
	CanTrade         bool `json:"canTrade"`
	CanWithdraw      bool `json:"canWithdraw"`
	CanDeposit       bool `json:"canDeposit"`
	UpdateTime       int  `json:"updateTime"`
	Balances         []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}
