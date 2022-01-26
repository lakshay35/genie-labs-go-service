package models

type MultiCurrencyPrices struct {
	USDPrice float64 `json:"usdPrice"`
	JPYPrice float64 `json:"jpyPrice"`
	ETHPrice float64 `json:"ethPrice"`
}
