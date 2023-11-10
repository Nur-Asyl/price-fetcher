package types

type PriceResp struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}
