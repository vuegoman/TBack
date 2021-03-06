package model

import "time"

type CoinData struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}
type Status struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Notice       interface{} `json:"notice"`
}
type Usd struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
type Quote struct {
	Usd Usd `json:"USD"`
}
type Atom struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         interface{} `json:"max_supply"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             Quote       `json:"quote"`
}
type Btc struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         int         `json:"max_supply"`
	CirculatingSupply int         `json:"circulating_supply"`
	TotalSupply       int         `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             Quote       `json:"quote"`
}
type Eth struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         interface{} `json:"max_supply"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             Quote       `json:"quote"`
}
type Ltc struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         int         `json:"max_supply"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       int         `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             Quote       `json:"quote"`
}
type Xrp struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         int64       `json:"max_supply"`
	CirculatingSupply int64       `json:"circulating_supply"`
	TotalSupply       int64       `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             Quote       `json:"quote"`
}
type Data struct {
	Atom Atom `json:"ATOM"`
	Btc  Btc  `json:"BTC"`
	Eth  Eth  `json:"ETH"`
	Ltc  Ltc  `json:"LTC"`
	Xrp  Xrp  `json:"XRP"`
}
