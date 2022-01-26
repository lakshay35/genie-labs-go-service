package models

type EthereumExchangeQuote struct {
	Data Data `json:"data"`
}

type Data struct {
	Currency string            `json:"currency"`
	Rates    map[string]string `json:"rates"`
}
