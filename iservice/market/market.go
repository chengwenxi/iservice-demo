package market

import (
	"iservice/iservice/market/bian"
	"iservice/iservice/market/huobi"
	"math/rand"
)

type Market interface {
	GetPrice(base string, quote string) (price float64, error string)
}

type RandomMarket struct {
}

func (market RandomMarket) GetPrice(base string, quote string) (price float64, error string) {
	return rand.Float64(), ""
}

var MarketType string

func GetMarket() Market {
	switch MarketType {
	case "huobi":
		return huobi.HuobiMarket{}
	case "binance":
		return bian.BianMarket{}
	default:
		return RandomMarket{}
	}
}
