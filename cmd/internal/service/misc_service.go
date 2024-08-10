package service

import (
	"strconv"
)

var DollarExchangeRateKey = "dollar_exchange_rate"

func GetCurrentDollarExchange() (float32, *APIError) {
	value, err := configRepo.Get(DollarExchangeRateKey)
	if err != nil {
		return 0, ErrorCurrentExchangeNotFound
	}

	float, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, ErrorCurrentExchangeNotFound
	}
	return float32(float), nil
}
